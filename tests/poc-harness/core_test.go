package poc

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type fixture struct {
	Classification string          `json:"classification"`
	Notice         string          `json:"notice"`
	Accounts       []account       `json:"accounts"`
	Profiles       []profile       `json:"profiles"`
	Relationships  []relationship  `json:"relationships"`
	ProviderEvents []providerEvent `json:"provider_events"`
	Scenarios      []scenario      `json:"scenarios"`
	Constraints    map[string]bool `json:"constraints"`
}
type account struct {
	ID      string `json:"id"`
	State   string `json:"state"`
	AgeGate string `json:"age_gate"`
	Profile string `json:"profile"`
}
type profile struct {
	ID               string   `json:"id"`
	AccountID        string   `json:"account_id"`
	RelationshipGoal string   `json:"relationship_goal"`
	OpenTo           []string `json:"open_to"`
}
type relationship struct {
	ID   string   `json:"id"`
	Type string   `json:"type"`
	From string   `json:"from"`
	To   string   `json:"to"`
	Pair []string `json:"pair"`
}
type providerEvent struct {
	ID              string `json:"id"`
	Provider        string `json:"provider"`
	AccountID       string `json:"account_id"`
	WorkflowVersion string `json:"workflow_version"`
	Signature       string `json:"signature"`
}
type scenario struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	Assert []string `json:"assert"`
}

func loadFixture(t *testing.T) fixture {
	t.Helper()
	path := filepath.Join("..", "fixtures", "poc", "core-scenarios.json")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read synthetic fixture: %v", err)
	}
	var f fixture
	if err := json.Unmarshal(raw, &f); err != nil {
		t.Fatalf("decode fixture JSON: %v", err)
	}
	return f
}

func TestFixtureIsExplicitlySyntheticAndOffline(t *testing.T) {
	f := loadFixture(t)
	if f.Classification != "synthetic-only" {
		t.Fatalf("classification = %q", f.Classification)
	}
	if !strings.Contains(strings.ToLower(f.Notice), "never submit to live providers") {
		t.Fatal("fixture lacks live-provider prohibition")
	}
	for _, key := range []string{"no_real_person_data", "no_live_provider_calls", "no_real_notifications", "no_production_credentials", "no_exact_real_world_coordinates", "fixture_generation_is_not_test_evidence"} {
		if f.Constraints[key] != true {
			t.Errorf("constraint %q must be true", key)
		}
	}
	for _, e := range f.ProviderEvents {
		if e.Provider != "age_assurance_fake" || e.WorkflowVersion != "local-fake-v1" || e.Signature != "LOCAL_TEST_ONLY" {
			t.Errorf("provider event %q is not an explicitly fake local event", e.ID)
		}
	}
}

func TestFixtureReferencesAndIdentifiersAreConsistent(t *testing.T) {
	f := loadFixture(t)
	accounts, profiles := map[string]account{}, map[string]profile{}
	ids := map[string]bool{}
	for _, a := range f.Accounts {
		if a.ID == "" || ids[a.ID] {
			t.Fatalf("empty or duplicate account id %q", a.ID)
		}
		ids[a.ID] = true
		accounts[a.ID] = a
	}
	for _, p := range f.Profiles {
		if p.ID == "" || ids[p.ID] {
			t.Fatalf("empty or duplicate profile id %q", p.ID)
		}
		ids[p.ID] = true
		profiles[p.ID] = p
		if _, ok := accounts[p.AccountID]; !ok {
			t.Errorf("profile %s references missing account %s", p.ID, p.AccountID)
		}
	}
	for _, a := range f.Accounts {
		p, ok := profiles[a.Profile]
		if !ok {
			t.Errorf("account %s references missing profile %s", a.ID, a.Profile)
			continue
		}
		if p.AccountID != a.ID {
			t.Errorf("account/profile reverse link mismatch for %s", a.ID)
		}
	}
	for _, r := range f.Relationships {
		if r.From != "" {
			if _, ok := accounts[r.From]; !ok {
				t.Errorf("relationship %s has missing from account", r.ID)
			}
		}
		if r.To != "" {
			if _, ok := accounts[r.To]; !ok {
				t.Errorf("relationship %s has missing to account", r.ID)
			}
		}
		for _, id := range r.Pair {
			if _, ok := accounts[id]; !ok {
				t.Errorf("relationship %s has missing pair account %s", r.ID, id)
			}
		}
	}
	for _, e := range f.ProviderEvents {
		if _, ok := accounts[e.AccountID]; !ok {
			t.Errorf("provider event %s references missing account", e.ID)
		}
	}
	if len(f.Scenarios) < 10 {
		t.Errorf("expected at least 10 named scenarios; got %d", len(f.Scenarios))
	}
}

func basicAccessAllowed(a account) bool {
	return a.State == "active" && a.AgeGate == "verified_test_state"
}

func TestAgeGateAndAccountStateAreIndependentHardGates(t *testing.T) {
	f := loadFixture(t)
	expected := map[string]bool{
		"acct-001": true,
		"acct-002": true,
		"acct-003": true,
		"acct-004": false, // paused
		"acct-005": false, // age assurance pending
		"acct-006": false, // deletion pending
		"acct-007": false, // age assurance expired
		"acct-008": false, // suspended
	}
	for _, a := range f.Accounts {
		want, ok := expected[a.ID]
		if !ok {
			t.Errorf("missing expected access result for %s", a.ID)
			continue
		}
		if got := basicAccessAllowed(a); got != want {
			t.Errorf("basicAccessAllowed(%s) = %v, want %v", a.ID, got, want)
		}
		delete(expected, a.ID)
	}
	for id := range expected {
		t.Errorf("expected account %s is absent from fixture", id)
	}
}

func TestUnknownRelationshipGoalIsNotPositiveIntent(t *testing.T) {
	f := loadFixture(t)
	found := false
	for _, p := range f.Profiles {
		if p.RelationshipGoal == "not_sure_yet" {
			found = true
			if len(p.OpenTo) != 0 {
				t.Errorf("profile %s has unknown goal but explicit open-to values; do not infer intent", p.ID)
			}
		}
	}
	if !found {
		t.Fatal("fixture no longer covers not_sure_yet")
	}
}

func TestBlockAndUnmatchSuppressionAreSeparateRelationshipTypes(t *testing.T) {
	f := loadFixture(t)
	seenBlock, seenSuppression := false, false
	for _, r := range f.Relationships {
		if r.Type == "block" {
			seenBlock = true
		}
		if r.Type == "unmatch_suppression" {
			seenSuppression = true
		}
	}
	if !seenBlock || !seenSuppression {
		t.Fatal("fixture must separately cover block and post-unmatch pair suppression")
	}
}

func TestScenarioIDsAreUniqueAndAssertionsPresent(t *testing.T) {
	f := loadFixture(t)
	seen := map[string]bool{}
	for _, s := range f.Scenarios {
		if s.ID == "" || seen[s.ID] {
			t.Errorf("empty or duplicate scenario id %q", s.ID)
		}
		seen[s.ID] = true
		if s.Name == "" || len(s.Assert) == 0 {
			t.Errorf("scenario %s needs a name and at least one assertion", s.ID)
		}
	}
}
