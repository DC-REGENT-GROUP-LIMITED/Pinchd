package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

func TestLoadConfigRequiresDatabaseCredentials(t *testing.T) {
	_, err := loadConfig(func(string) string { return "" })
	if err == nil || !strings.Contains(err.Error(), "configuration incomplete") {
		t.Fatalf("expected safe missing-config error, got %v", err)
	}
}

func TestLoadConfigRejectsInvalidPort(t *testing.T) {
	values := map[string]string{
		"PGHOST": "127.0.0.1", "PGDATABASE": "pinchd_poc", "PGUSER": "test",
		"PGPASSWORD": "not-a-real-secret", "PGPORT": "70000",
	}
	_, err := loadConfig(func(k string) string { return values[k] })
	if err == nil || !strings.Contains(err.Error(), "PGPORT") {
		t.Fatalf("expected port validation error, got %v", err)
	}
}

func TestHealthAndReadinessResponses(t *testing.T) {
	db, err := sql.Open("pgx", "postgres://invalid:invalid@127.0.0.1:1/invalid?connect_timeout=1&sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	h := (api{db: db}).routes()

	health := httptest.NewRecorder()
	h.ServeHTTP(health, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	if health.Code != http.StatusOK || !strings.Contains(health.Body.String(), `"status":"ok"`) {
		t.Fatalf("unexpected health response: %d %s", health.Code, health.Body.String())
	}
	if health.Header().Get("Cache-Control") != "no-store" || health.Header().Get("X-Content-Type-Options") != "nosniff" {
		t.Fatal("security headers missing")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	req := httptest.NewRequest(http.MethodGet, "/readyz", nil).WithContext(ctx)
	ready := httptest.NewRecorder()
	h.ServeHTTP(ready, req)
	if ready.Code != http.StatusServiceUnavailable || strings.Contains(ready.Body.String(), "invalid") || strings.Contains(ready.Body.String(), "127.0.0.1") {
		t.Fatalf("readiness must fail closed without leaking DB details: %d %s", ready.Code, ready.Body.String())
	}
}

func TestIntegrationDatabaseReadiness(t *testing.T) {
	if os.Getenv("PINCHD_POC_INTEGRATION") != "1" {
		t.Skip("set PINCHD_POC_INTEGRATION=1 to run against the isolated local PoC database")
	}
	c, err := loadConfig(os.Getenv)
	if err != nil {
		t.Fatal(err)
	}
	db, err := sql.Open("pgx", c.dataSourceName())
	if err != nil {
		t.Fatal("database client initialization failed")
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		t.Fatal("isolated database readiness integration failed")
	}
	response := httptest.NewRecorder()
	(api{db: db}).routes().ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/readyz", nil))
	if response.Code != http.StatusOK {
		t.Fatalf("readyz status = %d; expected 200", response.Code)
	}
	var body map[string]string
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil || body["status"] != "ready" {
		t.Fatalf("unexpected readiness body: %s", response.Body.String())
	}
}
