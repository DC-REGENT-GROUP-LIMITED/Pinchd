package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type config struct {
	ListenAddr string
	Host       string
	Port       int
	Database   string
	User       string
	Password   string
}

func loadConfig(getenv func(string) string) (config, error) {
	c := config{
		ListenAddr: strings.TrimSpace(getenv("PINCHD_LISTEN_ADDR")),
		Host:       strings.TrimSpace(getenv("PGHOST")),
		Database:   strings.TrimSpace(getenv("PGDATABASE")),
		User:       strings.TrimSpace(getenv("PGUSER")),
		Password:   getenv("PGPASSWORD"),
		Port:       5432,
	}
	if c.ListenAddr == "" {
		c.ListenAddr = "127.0.0.1:8080"
	}
	if c.Host == "" || c.Database == "" || c.User == "" || c.Password == "" {
		return config{}, errors.New("database configuration incomplete: set PGHOST, PGPORT (optional), PGDATABASE, PGUSER and PGPASSWORD")
	}
	if raw := strings.TrimSpace(getenv("PGPORT")); raw != "" {
		port, err := strconv.Atoi(raw)
		if err != nil || port < 1 || port > 65535 {
			return config{}, errors.New("PGPORT must be an integer between 1 and 65535")
		}
		c.Port = port
	}
	host, port, err := net.SplitHostPort(c.ListenAddr)
	if err != nil || strings.TrimSpace(host) == "" {
		return config{}, errors.New("PINCHD_LISTEN_ADDR must be a host:port address")
	}
	listenPort, err := strconv.Atoi(port)
	if err != nil || listenPort < 1 || listenPort > 65535 {
		return config{}, errors.New("PINCHD_LISTEN_ADDR port must be between 1 and 65535")
	}
	return c, nil
}

func (c config) dataSourceName() string {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(c.User, c.Password),
		Host:   net.JoinHostPort(c.Host, strconv.Itoa(c.Port)),
		Path:   "/" + c.Database,
	}
	q := url.Values{}
	q.Set("sslmode", "disable") // local loopback PoC only; production transport needs review.
	q.Set("connect_timeout", "3")
	u.RawQuery = q.Encode()
	return u.String()
}

type api struct {
	db *sql.DB
}

func (a api) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 1500*time.Millisecond)
		defer cancel()
		if err := a.db.PingContext(ctx); err != nil {
			// Do not return database, network or credential details to callers.
			writeJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "not_ready"})
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
	})
	return securityHeaders(mux)
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func run() error {
	c, err := loadConfig(os.Getenv)
	if err != nil {
		return err
	}
	db, err := sql.Open("pgx", c.dataSourceName())
	if err != nil {
		return errors.New("database client initialization failed")
	}
	defer db.Close()
	db.SetMaxOpenConns(8)
	db.SetMaxIdleConns(2)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(30 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		// Log a generic startup error only; never print DSN or driver error.
		return errors.New("database is unavailable; check local configuration and database status")
	}

	server := &http.Server{
		Addr:              c.ListenAddr,
		Handler:           (api{db: db}).routes(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	errCh := make(chan error, 1)
	go func() {
		errCh <- server.ListenAndServe()
	}()

	stopCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	select {
	case <-stopCtx.Done():
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			return errors.New("HTTP server shutdown timed out")
		}
		return nil
	case err := <-errCh:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return fmt.Errorf("HTTP server failed to start: %w", err)
	}
}

func main() {
	if err := run(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}
