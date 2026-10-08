package config

import (
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/VictoriaMetrics/VictoriaMetrics/lib/auth"
	vmcloud "github.com/VictoriaMetrics/victoriametrics-cloud-api-go/v1"
	"gopkg.in/yaml.v2"
)

const (
	toolsDisabledByDefault = "export,flags,metric_relabel_debug,downsampling_filters_debug,retention_filters_debug,test_rules"
)

var tenantAliasRegexp = regexp.MustCompile(`^[a-z][a-z0-9_-]*$`)

// Tenant describes the tenant of VictoriaMetrics cluster in VM_TENANTS_FILE
type Tenant struct {
	// ID is the tenant ID in accountID:projectID format
	ID string `yaml:"id"`
	// Alias is the short name of the tenant (e.g. environment name), which can be used in tools instead of the tenant ID
	Alias string `yaml:"alias,omitempty"`
	// Description is the free-form description of the tenant, which helps AI assistant to choose the right tenant
	Description string `yaml:"description,omitempty"`
}

type Config struct {
	serverMode         string
	listenAddr         string
	entrypoint         string
	instanceType       string
	bearerToken        string
	disabledTools      map[string]bool
	apiKey             string
	apiBaseURL         string
	heartbeatInterval  time.Duration
	disableResources   bool
	customHeaders      map[string]string
	passthroughHeaders []string
	defaultTenantID    string
	tenants            []Tenant

	// Logging configuration
	logFormat string
	logLevel  string

	entryPointURL *url.URL
	vmc           *vmcloud.VMCloudAPIClient
}

func InitConfig() (*Config, error) {
	disabledTools, isDisabledToolsSet := os.LookupEnv("MCP_DISABLED_TOOLS")
	if disabledTools == "" && !isDisabledToolsSet {
		disabledTools = toolsDisabledByDefault
	}
	disabledToolsMap := make(map[string]bool)
	if disabledTools != "" {
		for _, tool := range strings.Split(disabledTools, ",") {
			tool = strings.Trim(tool, " ,")
			if tool != "" {
				disabledToolsMap[tool] = true
			}
		}
	}

	customHeaders := os.Getenv("VM_INSTANCE_HEADERS")
	customHeadersMap := make(map[string]string)
	if customHeaders != "" {
		for _, header := range strings.Split(customHeaders, ",") {
			header = strings.TrimSpace(header)
			if header != "" {
				parts := strings.SplitN(header, "=", 2)
				if len(parts) == 2 {
					key := strings.TrimSpace(parts[0])
					value := strings.TrimSpace(parts[1])
					if key != "" && value != "" {
						customHeadersMap[key] = value
					}
				}
			}
		}
	}

	var passthroughHeaders []string
	passthroughHeadersStr := os.Getenv("MCP_PASSTHROUGH_HEADERS")
	if passthroughHeadersStr != "" {
		for _, h := range strings.Split(passthroughHeadersStr, ",") {
			h = strings.TrimSpace(h)
			if h != "" {
				passthroughHeaders = append(passthroughHeaders, h)
			}
		}
	}

	heartbeatInterval := 30 * time.Second
	heartbeatIntervalStr := os.Getenv("MCP_HEARTBEAT_INTERVAL")
	if heartbeatIntervalStr != "" {
		interval, err := time.ParseDuration(heartbeatIntervalStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse MCP_HEARTBEAT_INTERVAL: %w", err)
		}
		if interval < 0 {
			return nil, fmt.Errorf("MCP_HEARTBEAT_INTERVAL must be a non-negative")
		}
		heartbeatInterval = interval
	}

	disableResources := false
	disableResourcesStr := os.Getenv("MCP_DISABLE_RESOURCES")
	if disableResourcesStr != "" {
		var err error
		disableResources, err = strconv.ParseBool(disableResourcesStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse MCP_DISABLE_RESOURCES: %w", err)
		}
	}

	logFormat := strings.ToLower(os.Getenv("MCP_LOG_FORMAT"))
	if logFormat == "" {
		logFormat = "text"
	}
	if logFormat != "text" && logFormat != "json" {
		return nil, fmt.Errorf("MCP_LOG_FORMAT must be 'text' or 'json'")
	}

	logLevel := strings.ToLower(os.Getenv("MCP_LOG_LEVEL"))
	if logLevel == "" {
		logLevel = "info"
	}
	if logLevel != "debug" && logLevel != "info" && logLevel != "warn" && logLevel != "error" {
		return nil, fmt.Errorf("MCP_LOG_LEVEL must be 'debug', 'info', 'warn' or 'error'")
	}
	result := &Config{
		serverMode:         strings.ToLower(os.Getenv("MCP_SERVER_MODE")),
		listenAddr:         os.Getenv("MCP_LISTEN_ADDR"),
		entrypoint:         os.Getenv("VM_INSTANCE_ENTRYPOINT"),
		instanceType:       os.Getenv("VM_INSTANCE_TYPE"),
		bearerToken:        os.Getenv("VM_INSTANCE_BEARER_TOKEN"),
		disabledTools:      disabledToolsMap,
		apiKey:             os.Getenv("VMC_API_KEY"),
		apiBaseURL:         os.Getenv("VMC_API_BASE_URL"),
		heartbeatInterval:  heartbeatInterval,
		disableResources:   disableResources,
		customHeaders:      customHeadersMap,
		passthroughHeaders: passthroughHeaders,
		logFormat:          logFormat,
		logLevel:           logLevel,
		defaultTenantID:    "0",
	}
	// Left for backward compatibility
	if result.listenAddr == "" {
		result.listenAddr = os.Getenv("MCP_SSE_ADDR")
	}
	if result.entrypoint == "" && result.apiKey == "" {
		return nil, fmt.Errorf("VM_INSTANCE_ENTRYPOINT or VMC_API_KEY is not set")
	}
	if result.entrypoint != "" && result.apiKey != "" {
		return nil, fmt.Errorf("VM_INSTANCE_ENTRYPOINT and VMC_API_KEY cannot be set at the same time")
	}
	if result.entrypoint != "" && result.instanceType == "" {
		return nil, fmt.Errorf("VM_INSTANCE_TYPE is not set")
	}
	if result.entrypoint != "" && result.instanceType != "cluster" && result.instanceType != "single" {
		return nil, fmt.Errorf("VM_INSTANCE_TYPE must be 'single' or 'cluster'")
	}
	if result.serverMode != "" && result.serverMode != "stdio" && result.serverMode != "sse" && result.serverMode != "http" {
		return nil, fmt.Errorf("MCP_SERVER_MODE must be 'stdio', 'sse' or 'http'")
	}
	if result.serverMode == "" {
		result.serverMode = "stdio"
	}
	if result.listenAddr == "" {
		result.listenAddr = "localhost:8080"
	}

	var err error
	if result.apiKey == "" {
		result.entryPointURL, err = url.Parse(result.entrypoint)
		if err != nil {
			return nil, fmt.Errorf("failed to parse URL from VM_INSTANCE_ENTRYPOINT: %w", err)
		}
	}
	if result.apiKey != "" {
		vmcOptions := []vmcloud.VMCloudAPIClientOption{}
		if result.apiBaseURL != "" {
			vmcOptions = append(vmcOptions, vmcloud.WithBaseURL(result.apiBaseURL))
		}
		result.vmc, err = vmcloud.New(result.apiKey, vmcOptions...)
		if err != nil {
			return nil, fmt.Errorf("failed to create VMCloud API client: %w", err)
		}
	}

	tenantsFile := os.Getenv("VM_TENANTS_FILE")
	if tenantsFile != "" {
		if !result.IsCluster() {
			return nil, fmt.Errorf("VM_TENANTS_FILE can be used only with VM_INSTANCE_TYPE=cluster")
		}
		result.tenants, err = loadTenants(tenantsFile)
		if err != nil {
			return nil, err
		}
	}

	defaultTenantID := strings.ToLower(os.Getenv("VM_DEFAULT_TENANT_ID"))
	if t, ok := result.TenantByAlias(defaultTenantID); ok {
		result.defaultTenantID = t.ID
	} else if defaultTenantID != "" {
		tenantID, err := auth.NewToken(defaultTenantID)
		if err != nil {
			return nil, fmt.Errorf("failed to parse VM_DEFAULT_TENANT_ID %q: %w", defaultTenantID, err)
		}
		result.defaultTenantID = tenantID.String()
	}

	return result, nil
}

func loadTenants(path string) ([]Tenant, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read VM_TENANTS_FILE: %w", err)
	}
	var file struct {
		Tenants []Tenant `yaml:"tenants"`
	}
	if err := yaml.UnmarshalStrict(data, &file); err != nil {
		return nil, fmt.Errorf("failed to parse VM_TENANTS_FILE %q: %w", path, err)
	}
	if len(file.Tenants) == 0 {
		return nil, fmt.Errorf("VM_TENANTS_FILE %q doesn't contain any tenants", path)
	}

	ids := make(map[string]bool)
	aliases := make(map[string]bool)
	for i := range file.Tenants {
		t := &file.Tenants[i]
		if t.ID == "" {
			return nil, fmt.Errorf("tenant #%d in VM_TENANTS_FILE has no id", i+1)
		}
		id, err := CanonicalTenantID(t.ID)
		if err != nil {
			return nil, fmt.Errorf("invalid tenant ID %q in VM_TENANTS_FILE: %w", t.ID, err)
		}
		if ids[id] {
			return nil, fmt.Errorf("duplicate tenant ID %q in VM_TENANTS_FILE", id)
		}
		ids[id] = true
		t.ID = id

		if t.Alias != "" {
			if !tenantAliasRegexp.MatchString(t.Alias) {
				return nil, fmt.Errorf("invalid alias %q of tenant %q in VM_TENANTS_FILE: it must start with a lowercase letter and contain only lowercase letters, digits, '_' and '-'", t.Alias, id)
			}
			if aliases[t.Alias] {
				return nil, fmt.Errorf("duplicate tenant alias %q in VM_TENANTS_FILE", t.Alias)
			}
			aliases[t.Alias] = true
		}
		t.Description = strings.TrimSpace(t.Description)
	}
	return file.Tenants, nil
}

// CanonicalTenantID returns the given tenant ID in accountID:projectID format (the same as /admin/tenants API returns)
func CanonicalTenantID(id string) (string, error) {
	t, err := auth.NewToken(id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d:%d", t.AccountID, t.ProjectID), nil
}

func (c *Config) IsCluster() bool {
	return c.instanceType == "cluster"
}

func (c *Config) IsSingle() bool {
	return c.instanceType == "single"
}

func (c *Config) IsStdio() bool {
	return c.serverMode == "stdio"
}

func (c *Config) IsSSE() bool {
	return c.serverMode == "sse"
}

func (c *Config) ServerMode() string {
	return c.serverMode
}

func (c *Config) IsCloud() bool {
	return c.vmc != nil
}

func (c *Config) IsCloudSharedInstance() bool {
	return c.apiKey == vmcloud.DynamicAPIKey
}

func (c *Config) VMC() *vmcloud.VMCloudAPIClient {
	return c.vmc
}

func (c *Config) ListenAddr() string {
	return c.listenAddr
}

func (c *Config) BearerToken() string {
	return c.bearerToken
}

func (c *Config) EntryPointURL() *url.URL {
	return c.entryPointURL
}

func (c *Config) IsToolDisabled(toolName string) bool {
	if c.disabledTools == nil {
		return false
	}
	disabled, ok := c.disabledTools[toolName]
	return ok && disabled
}

func (c *Config) IsResourcesDisabled() bool {
	return c.disableResources
}

func (c *Config) HeartbeatInterval() time.Duration {
	return c.heartbeatInterval
}

func (c *Config) CustomHeaders() map[string]string {
	return c.customHeaders
}

func (c *Config) PassthroughHeaders() []string {
	return c.passthroughHeaders
}

func (c *Config) LogFormat() string {
	return c.logFormat
}

func (c *Config) LogLevel() string {
	return c.logLevel
}

func (c *Config) DefaultTenantID() string {
	return c.defaultTenantID
}

// Tenants returns the tenants described in VM_TENANTS_FILE
func (c *Config) Tenants() []Tenant {
	return c.tenants
}

// TenantByAlias returns the tenant from VM_TENANTS_FILE with the given alias (case-insensitive)
func (c *Config) TenantByAlias(alias string) (Tenant, bool) {
	alias = strings.ToLower(alias)
	for _, t := range c.tenants {
		if t.Alias != "" && t.Alias == alias {
			return t, true
		}
	}
	return Tenant{}, false
}
