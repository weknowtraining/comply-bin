package jira

import (
	"os"
	"testing"
)

func TestConfigureExpandsPasswordEnvVar(t *testing.T) {
	t.Setenv("JIRA_PASSWORD", "token-from-env")

	j := &jiraPlugin{}
	err := j.Configure(map[string]interface{}{
		cfgUsername: "user@example.com",
		cfgPassword: "${JIRA_PASSWORD}",
		cfgURL:      "https://example.atlassian.net",
		cfgProject:  "ABC",
		cfgTaskType: "Task",
	})
	if err != nil {
		t.Fatalf("Configure returned error: %v", err)
	}
	if j.password != "token-from-env" {
		t.Fatalf("expected expanded password, got %q", j.password)
	}
}

func TestConfigureKeepsLiteralPassword(t *testing.T) {
	os.Unsetenv("JIRA_PASSWORD")

	j := &jiraPlugin{}
	err := j.Configure(map[string]interface{}{
		cfgUsername: "user@example.com",
		cfgPassword: "literal-secret",
		cfgURL:      "https://example.atlassian.net",
		cfgProject:  "ABC",
		cfgTaskType: "Task",
	})
	if err != nil {
		t.Fatalf("Configure returned error: %v", err)
	}
	if j.password != "literal-secret" {
		t.Fatalf("expected literal password, got %q", j.password)
	}
}
