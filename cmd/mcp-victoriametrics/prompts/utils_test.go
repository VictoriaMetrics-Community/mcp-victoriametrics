package prompts

import (
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestGetPromptReqParam(t *testing.T) {
	// Test cases
	testCases := []struct {
		name          string
		args          map[string]string
		param         string
		required      bool
		expectedValue string
		expectError   bool
	}{
		{
			name:          "Valid parameter",
			args:          map[string]string{"test": "value"},
			param:         "test",
			required:      true,
			expectedValue: "value",
			expectError:   false,
		},
		{
			name:          "Missing required parameter",
			args:          map[string]string{},
			param:         "test",
			required:      true,
			expectedValue: "",
			expectError:   true,
		},
		{
			name:          "Missing optional parameter",
			args:          map[string]string{},
			param:         "test",
			required:      false,
			expectedValue: "",
			expectError:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock prompt request
			gpr := mcp.GetPromptRequest{}
			gpr.Params.Arguments = tc.args

			// Call the function
			value, err := GetPromptReqParam(gpr, tc.param, tc.required)

			// Check the result
			if tc.expectError && err == nil {
				t.Error("Expected an error, got nil")
			}
			if !tc.expectError && err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}
			if value != tc.expectedValue {
				t.Errorf("Expected '%s', got: '%s'", tc.expectedValue, value)
			}
		})
	}
}
