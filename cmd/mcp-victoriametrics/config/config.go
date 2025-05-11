package config

import (
	"fmt"
	"net/url"
	"os"
)

type Config struct {
	serverMode   string
	sseAddr      string
	entrypoint   string
	instanceType string
	bearerToken  string

	entryPointURL *url.URL
}

func InitConfig() (*Config, error) {
	result := &Config{
		serverMode:   os.Getenv("MCP_SERVER_MODE"),
		sseAddr:      os.Getenv("MCP_SSE_ADDR"),
		entrypoint:   os.Getenv("VM_INSTANCE_ENTRYPOINT"),
		instanceType: os.Getenv("VM_INSTANCE_TYPE"),
		bearerToken:  os.Getenv("VM_INSTANCE_BEARER_TOKEN"),
	}
	if result.entrypoint == "" {
		return nil, fmt.Errorf("VM_INSTANCE_ENTRYPOINT is not set")
	}
	if result.instanceType == "" {
		return nil, fmt.Errorf("VM_INSTANCE_TYPE is not set")
	}
	if result.instanceType != "cluster" && result.instanceType != "single" {
		return nil, fmt.Errorf("VM_INSTANCE_TYPE must be 'single' or 'cluster'")
	}
	if result.serverMode != "" && result.serverMode != "stdio" && result.serverMode != "sse" {
		return nil, fmt.Errorf("MCP_SERVER_MODE must be 'stdio' or 'sse'")
	}
	if result.serverMode == "" {
		result.serverMode = "stdio"
	}
	if result.sseAddr == "" {
		result.sseAddr = "localhost:8080"
	}

	var err error
	result.entryPointURL, err = url.Parse(result.entrypoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL from VM_INSTANCE_ENTRYPOINT: %w", err)
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

func (c *Config) SSEAddr() string {
	return c.sseAddr
}

func (c *Config) BearerToken() string {
	return c.bearerToken
}

func (c *Config) EntryPointURL() *url.URL {
	return c.entryPointURL
}

func (c *Config) AdminAPIURL(path ...string) string {
	return c.entryPointURL.JoinPath(path...).String()
}

func (c *Config) SelectAPIURL(tenant string, path ...string) string {
	if c.IsSingle() {
		return c.entryPointURL.JoinPath(path...).String()
	}
	if tenant == "" {
		tenant = "0"
	}
	args := []string{"select", tenant, "prometheus"}
	return c.entryPointURL.JoinPath(append(args, path...)...).String()
}
