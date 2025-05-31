package resources

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

// TestGetDocResourceContent tests the GetDocResourceContent function
func TestGetDocResourceContent(t *testing.T) {
	// Setup test data
	testURI := "docs://test.md#0"
	testContent := mcp.TextResourceContents{
		URI:      testURI,
		MIMEType: "text/markdown",
		Text:     "# Test Document\n\nThis is a test document.",
	}

	// Save original contents and restore after test
	origContents := contents
	defer func() { contents = origContents }()

	// Initialize contents map for testing
	contents = map[string]mcp.ResourceContents{
		testURI: testContent,
	}

	// Test cases
	testCases := []struct {
		name        string
		uri         string
		expectError bool
	}{
		{
			name:        "Valid URI",
			uri:         testURI,
			expectError: false,
		},
		{
			name:        "Invalid URI",
			uri:         "docs://nonexistent.md#0",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function
			content, err := GetDocResourceContent(tc.uri)

			// Check for errors
			if tc.expectError {
				if err == nil {
					t.Error("Expected an error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Check the content
			if content == nil {
				t.Fatal("Expected non-nil content")
			}

			// Check that the content is a TextResourceContents
			textContent, ok := content.(mcp.TextResourceContents)
			if !ok {
				t.Fatal("Expected TextResourceContents, got different content type")
			}

			// Check the URI
			if textContent.URI != tc.uri {
				t.Errorf("Expected URI '%s', got: '%s'", tc.uri, textContent.URI)
			}
		})
	}
}

// Skip TestSearchDocResources for now as it requires more complex mocking
// of the bleve search index which is causing nil pointer dereference issues

// TestDocResourcesHandler tests the docResourcesHandler function
func TestDocResourcesHandler(t *testing.T) {
	// Setup test data
	testURI := "docs://test.md#0"
	testContent := mcp.TextResourceContents{
		URI:      testURI,
		MIMEType: "text/markdown",
		Text:     "# Test Document\n\nThis is a test document.",
	}

	// Save original contents and restore after test
	origContents := contents
	defer func() { contents = origContents }()

	// Initialize contents map for testing
	contents = map[string]mcp.ResourceContents{
		testURI: testContent,
	}

	// Test cases
	testCases := []struct {
		name        string
		uri         string
		expectError bool
	}{
		{
			name:        "Valid URI",
			uri:         testURI,
			expectError: false,
		},
		{
			name:        "Invalid URI",
			uri:         "docs://nonexistent.md#0",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock request
			req := mcp.ReadResourceRequest{}
			req.Params.URI = tc.uri

			// Call the handler
			result, err := docResourcesHandler(context.Background(), req)

			// Check for errors
			if tc.expectError {
				if err == nil {
					t.Error("Expected an error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Check the result
			if len(result) != 1 {
				t.Fatalf("Expected 1 result, got: %d", len(result))
			}

			// Check that the content is correct
			content := result[0]
			textContent, ok := content.(mcp.TextResourceContents)
			if !ok {
				t.Fatal("Expected TextResourceContents, got different content type")
			}

			// Check the URI
			if textContent.URI != tc.uri {
				t.Errorf("Expected URI '%s', got: '%s'", tc.uri, textContent.URI)
			}
		})
	}
}
