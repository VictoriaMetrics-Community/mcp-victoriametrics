package tools

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/VictoriaMetrics/mcp-victoriametrics/cmd/mcp-victoriametrics/config"
)

const testTenantsFile = `
tenants:
  - id: "0:0"
    alias: prod
    description: Production environment
  - id: "1:0"
    alias: stage
  - id: "2:0"
    description: Load testing
`

// newTenantsTestConfig creates cluster config with the given content of VM_TENANTS_FILE (empty content means no file)
func newTenantsTestConfig(t *testing.T, entrypoint, tenantsFile, defaultTenantID string) *config.Config {
	t.Helper()
	path := ""
	if tenantsFile != "" {
		path = filepath.Join(t.TempDir(), "tenants.yaml")
		if err := os.WriteFile(path, []byte(tenantsFile), 0o600); err != nil {
			t.Fatalf("failed to write tenants file: %v", err)
		}
	}
	t.Setenv("VM_INSTANCE_ENTRYPOINT", entrypoint)
	t.Setenv("VM_INSTANCE_TYPE", "cluster")
	t.Setenv("VMC_API_KEY", "")
	t.Setenv("VM_TENANTS_FILE", path)
	t.Setenv("VM_DEFAULT_TENANT_ID", defaultTenantID)
	cfg, err := config.InitConfig()
	if err != nil {
		t.Fatalf("Failed to create config: %v", err)
	}
	return cfg
}

func TestGetSelectURLWithTenantAliases(t *testing.T) {
	testCases := []struct {
		name            string
		defaultTenantID string
		requestTenant   string
		expectedURL     string
		expectedErr     string
	}{
		{
			name:          "Alias is resolved to tenant ID",
			requestTenant: "stage",
			expectedURL:   "http://example.com/select/1:0/prometheus/api/v1/query",
		},
		{
			name:          "Alias is case-insensitive",
			requestTenant: "PROD",
			expectedURL:   "http://example.com/select/0:0/prometheus/api/v1/query",
		},
		{
			name:          "Tenant ID is passed as is",
			requestTenant: "5:1",
			expectedURL:   "http://example.com/select/5:1/prometheus/api/v1/query",
		},
		{
			name:            "Default tenant set by alias",
			defaultTenantID: "stage",
			expectedURL:     "http://example.com/select/1:0/prometheus/api/v1/query",
		},
		{
			name:          "Unknown alias",
			requestTenant: "qa",
			expectedErr:   `unknown tenant "qa": use tenant ID in accountID[:projectID] format or one of the aliases: prod (0:0), stage (1:0)`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := newTenantsTestConfig(t, "http://example.com", testTenantsFile, tc.defaultTenantID)

			tcr := mcp.CallToolRequest{}
			if tc.requestTenant != "" {
				tcr.Params.Arguments = map[string]any{"tenant": tc.requestTenant}
			}

			url, err := getSelectURL(context.Background(), cfg, tcr, "api", "v1", "query")
			if tc.expectedErr != "" {
				if err == nil || err.Error() != tc.expectedErr {
					t.Fatalf("Expected error %q, got: %v", tc.expectedErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if url != tc.expectedURL {
				t.Errorf("Expected URL %q, got %q", tc.expectedURL, url)
			}
		})
	}
}

func TestGetSelectURLWithoutTenantsFile(t *testing.T) {
	cfg := newTenantsTestConfig(t, "http://example.com", "", "")

	// Without VM_TENANTS_FILE the tenant parameter is passed to VictoriaMetrics as is
	tcr := mcp.CallToolRequest{}
	tcr.Params.Arguments = map[string]any{"tenant": "prod"}
	url, err := getSelectURL(context.Background(), cfg, tcr, "api", "v1", "query")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if expected := "http://example.com/select/prod/prometheus/api/v1/query"; url != expected {
		t.Errorf("Expected URL %q, got %q", expected, url)
	}
}

func TestToolTenantsHandler(t *testing.T) {
	const tenantsResponse = `{"status":"success","data":["0:0","1:0","3:0"]}`
	status := http.StatusOK
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/admin/tenants" {
			t.Errorf("Unexpected request path: %s", r.URL.Path)
		}
		w.WriteHeader(status)
		_, _ = w.Write([]byte(tenantsResponse))
	}))
	defer srv.Close()

	textContents := func(result *mcp.CallToolResult) []string {
		var texts []string
		for _, c := range result.Content {
			texts = append(texts, c.(mcp.TextContent).Text)
		}
		return texts
	}

	t.Run("Without tenants file", func(t *testing.T) {
		cfg := newTenantsTestConfig(t, srv.URL, "", "")
		result, _ := toolTenantsHandler(context.Background(), cfg, mcp.CallToolRequest{})
		texts := textContents(result)
		if result.IsError || len(texts) != 1 || texts[0] != tenantsResponse {
			t.Errorf("Expected unchanged response, got: %v", texts)
		}
	})

	t.Run("With tenants file", func(t *testing.T) {
		cfg := newTenantsTestConfig(t, srv.URL, testTenantsFile, "")
		result, _ := toolTenantsHandler(context.Background(), cfg, mcp.CallToolRequest{})
		texts := textContents(result)
		if result.IsError || len(texts) != 2 || texts[0] != tenantsResponse {
			t.Fatalf("Expected response with tenants description, got: %v", texts)
		}
		for _, line := range []string{
			"- 0:0 (alias: prod): Production environment\n",
			"- 1:0 (alias: stage)\n",
			"- 2:0: Load testing\n",
			`If the "tenant" parameter is omitted, tenant 0:0 (alias: prod) is used.`,
		} {
			if !strings.Contains(texts[1], line) {
				t.Errorf("Expected %q in tenants description, got: %s", line, texts[1])
			}
		}
	})

	t.Run("Described tenants are returned on API error", func(t *testing.T) {
		status = http.StatusForbidden
		defer func() { status = http.StatusOK }()

		cfg := newTenantsTestConfig(t, srv.URL, testTenantsFile, "")
		result, _ := toolTenantsHandler(context.Background(), cfg, mcp.CallToolRequest{})
		texts := textContents(result)
		if !result.IsError || len(texts) != 2 || !strings.Contains(texts[1], "- 1:0 (alias: stage)") {
			t.Errorf("Expected error with tenants description, got: %v", texts)
		}
	})
}

func TestTenantsInstructions(t *testing.T) {
	if instructions := TenantsInstructions(newTenantsTestConfig(t, "http://example.com", "", "")); instructions != "" {
		t.Errorf("Expected no instructions without tenants file, got: %q", instructions)
	}

	instructions := TenantsInstructions(newTenantsTestConfig(t, "http://example.com", testTenantsFile, "5"))
	for _, line := range []string{
		"- 0:0 (alias: prod): Production environment\n",
		`If the "tenant" parameter is omitted, tenant 5:0 is used.`,
		"ask the user",
	} {
		if !strings.Contains(instructions, line) {
			t.Errorf("Expected %q in instructions, got: %s", line, instructions)
		}
	}
}
