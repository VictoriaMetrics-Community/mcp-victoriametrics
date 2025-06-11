package config

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	vmcloud "github.com/VictoriaMetrics/victoriametrics-cloud-api-go/v1"
)

type Config struct {
	serverMode    string
	listenAddr    string
	entrypoint    string
	instanceType  string
	bearerToken   string
	disabledTools map[string]bool
	apiKey        string

	entryPointURL *url.URL
	vmc           *vmcloud.VMCloudAPIClient
}

func InitConfig() (*Config, error) {
	disabledTools := os.Getenv("MCP_DISABLED_TOOLS")
	disabledToolsMap := make(map[string]bool)
	if disabledTools != "" {
		for _, tool := range strings.Split(disabledTools, ",") {
			tool = strings.Trim(tool, " ,")
			if tool != "" {
				disabledToolsMap[tool] = true
			}
		}
	}
	result := &Config{
		serverMode:    strings.ToLower(os.Getenv("MCP_SERVER_MODE")),
		listenAddr:    os.Getenv("MCP_LISTEN_ADDR"),
		entrypoint:    os.Getenv("VM_INSTANCE_ENTRYPOINT"),
		instanceType:  os.Getenv("VM_INSTANCE_TYPE"),
		bearerToken:   os.Getenv("VM_INSTANCE_BEARER_TOKEN"),
		disabledTools: disabledToolsMap,
		apiKey:        os.Getenv("VMC_API_KEY"),
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
		result.vmc, err = vmcloud.New(result.apiKey)
		if err != nil {
			return nil, fmt.Errorf("failed to create VMCloud API client: %w", err)
		}
	}

	return result, nil
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
