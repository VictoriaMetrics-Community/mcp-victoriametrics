package config

import (
	"net/url"
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {
	// Save original environment variables
	originalEntrypoint := os.Getenv("VM_INSTANCE_ENTRYPOINT")
	originalInstanceType := os.Getenv("VM_INSTANCE_TYPE")
	originalServerMode := os.Getenv("MCP_SERVER_MODE")
	originalSSEAddr := os.Getenv("MCP_SSE_ADDR")
	originalBearerToken := os.Getenv("VM_INSTANCE_BEARER_TOKEN")

	// Restore environment variables after test
	defer func() {
		os.Setenv("VM_INSTANCE_ENTRYPOINT", originalEntrypoint)
		os.Setenv("VM_INSTANCE_TYPE", originalInstanceType)
		os.Setenv("MCP_SERVER_MODE", originalServerMode)
		os.Setenv("MCP_SSE_ADDR", originalSSEAddr)
		os.Setenv("VM_INSTANCE_BEARER_TOKEN", originalBearerToken)
	}()

	// Test case 1: Valid configuration
	t.Run("Valid configuration", func(t *testing.T) {
		// Set environment variables
		os.Setenv("VM_INSTANCE_ENTRYPOINT", "http://example.com")
		os.Setenv("VM_INSTANCE_TYPE", "single")
		os.Setenv("MCP_SERVER_MODE", "stdio")
		os.Setenv("MCP_SSE_ADDR", "localhost:8080")
		os.Setenv("VM_INSTANCE_BEARER_TOKEN", "test-token")

		// Initialize config
		cfg, err := InitConfig()

		// Check for errors
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		// Check config values
		if cfg.BearerToken() != "test-token" {
			t.Errorf("Expected bearer token 'test-token', got: %s", cfg.BearerToken())
		}
		if !cfg.IsSingle() {
			t.Error("Expected IsSingle() to be true")
		}
		if cfg.IsCluster() {
			t.Error("Expected IsCluster() to be false")
		}
		if !cfg.IsStdio() {
			t.Error("Expected IsStdio() to be true")
		}
		if cfg.IsSSE() {
			t.Error("Expected IsSSE() to be false")
		}
		if cfg.ListenAddr() != "localhost:8080" {
			t.Errorf("Expected SSE address 'localhost:8080', got: %s", cfg.ListenAddr())
		}
		expectedURL, _ := url.Parse("http://example.com")
		if cfg.EntryPointURL().String() != expectedURL.String() {
			t.Errorf("Expected entrypoint URL 'http://example.com', got: %s", cfg.EntryPointURL().String())
		}
		if !cfg.IsSingle() {
			t.Error("Expected IsSingle() to be true")
		}
		if cfg.IsCluster() {
			t.Error("Expected IsCluster() to be false")
		}
	})

	// Test case 2: Missing entrypoint
	t.Run("Missing entrypoint", func(t *testing.T) {
		// Set environment variables
		os.Setenv("VM_INSTANCE_ENTRYPOINT", "")
		os.Setenv("VM_INSTANCE_TYPE", "single")

		// Initialize config
		_, err := InitConfig()

		// Check for errors
		if err == nil {
			t.Fatal("Expected error for missing entrypoint, got nil")
		}
	})

	// Test case 3: Missing instance type
	t.Run("Missing instance type", func(t *testing.T) {
		// Set environment variables
		os.Setenv("VM_INSTANCE_ENTRYPOINT", "http://example.com")
		os.Setenv("VM_INSTANCE_TYPE", "")

		// Initialize config
		_, err := InitConfig()

		// Check for errors
		if err == nil {
			t.Fatal("Expected error for missing instance type, got nil")
		}
	})

	// Test case 4: Invalid instance type
	t.Run("Invalid instance type", func(t *testing.T) {
		// Set environment variables
		os.Setenv("VM_INSTANCE_ENTRYPOINT", "http://example.com")
		os.Setenv("VM_INSTANCE_TYPE", "invalid")

		// Initialize config
		_, err := InitConfig()

		// Check for errors
		if err == nil {
			t.Fatal("Expected error for invalid instance type, got nil")
		}
	})

	// Test case 5: Invalid server mode
	t.Run("Invalid server mode", func(t *testing.T) {
		// Set environment variables
		os.Setenv("VM_INSTANCE_ENTRYPOINT", "http://example.com")
		os.Setenv("VM_INSTANCE_TYPE", "single")
		os.Setenv("MCP_SERVER_MODE", "invalid")

		// Initialize config
		_, err := InitConfig()

		// Check for errors
		if err == nil {
			t.Fatal("Expected error for invalid server mode, got nil")
		}
	})

	// Test case 6: Default values
	t.Run("Default values", func(t *testing.T) {
		// Set environment variables
		os.Setenv("VM_INSTANCE_ENTRYPOINT", "http://example.com")
		os.Setenv("VM_INSTANCE_TYPE", "single")
		os.Setenv("MCP_SERVER_MODE", "")
		os.Setenv("MCP_SSE_ADDR", "")

		// Initialize config
		cfg, err := InitConfig()

		// Check for errors
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		// Check default values
		if !cfg.IsStdio() {
			t.Error("Expected default server mode to be stdio")
		}
		if cfg.ListenAddr() != "localhost:8080" {
			t.Errorf("Expected default SSE address 'localhost:8080', got: %s", cfg.ListenAddr())
		}
		if !cfg.IsSingle() {
			t.Error("Expected IsSingle() to be true")
		}
		if cfg.IsCluster() {
			t.Error("Expected IsCluster() to be false")
		}
	})

	// Test case 7: Cluster
	t.Run("Missing entrypoint", func(t *testing.T) {
		// Set environment variables
		os.Setenv("VM_INSTANCE_ENTRYPOINT", "http://example.com")
		os.Setenv("VM_INSTANCE_TYPE", "cluster")

		// Initialize config
		cfg, err := InitConfig()
		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		// Check values
		if cfg.IsSingle() {
			t.Error("Expected IsSingle() to be true")
		}
		if !cfg.IsCluster() {
			t.Error("Expected IsCluster() to be false")
		}
	})
}
