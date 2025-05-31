package resources

import (
	"strings"
	"testing"
)

func TestSplitMarkdown(t *testing.T) {
	// Test cases
	testCases := []struct {
		name           string
		content        string
		expectedChunks int
		checkFunc      func(t *testing.T, chunks []string)
	}{
		{
			name: "Simple markdown without front matter",
			content: `# Title
This is a simple markdown document.

## Section 1
Content of section 1.

## Section 2
Content of section 2.
`,
			expectedChunks: 3,
			checkFunc: func(t *testing.T, chunks []string) {
				if len(chunks) == 0 {
					t.Fatal("Expected at least one chunk")
				}

				// Check that the title is in one of the chunks
				titleFound := false
				section1Found := false
				section2Found := false

				for _, chunk := range chunks {
					if strings.Contains(chunk, "# Title") {
						titleFound = true
					}
					if strings.Contains(chunk, "## Section 1") {
						section1Found = true
					}
					if strings.Contains(chunk, "## Section 2") {
						section2Found = true
					}
				}

				if !titleFound {
					t.Error("Expected one of the chunks to contain the title")
				}
				if !section1Found {
					t.Error("Expected one of the chunks to contain Section 1")
				}
				if !section2Found {
					t.Error("Expected one of the chunks to contain Section 2")
				}
			},
		},
		{
			name: "Markdown with front matter",
			content: `---
title: "Document Title"
author: "Test Author"
---

This is a document with front matter.

## Section 1
Content of section 1.
`,
			expectedChunks: 2,
			checkFunc: func(t *testing.T, chunks []string) {
				if len(chunks) == 0 {
					t.Fatal("Expected at least one chunk")
				}

				// Check that the title and section are in the chunks
				titleFound := false
				section1Found := false

				for _, chunk := range chunks {
					if strings.Contains(chunk, "# Document Title") {
						titleFound = true
					}
					if strings.Contains(chunk, "## Section 1") {
						section1Found = true
					}
				}

				if !titleFound {
					t.Error("Expected one of the chunks to contain the title from front matter")
				}
				if !section1Found {
					t.Error("Expected one of the chunks to contain Section 1")
				}
			},
		},
		{
			name: "Markdown with front matter and space in title",
			content: `---
title : "Document Title With Space"
author: "Test Author"
---

This is a document with front matter.

## Section 1
Content of section 1.
`,
			expectedChunks: 2,
			checkFunc: func(t *testing.T, chunks []string) {
				if len(chunks) == 0 {
					t.Fatal("Expected at least one chunk")
				}

				// Check that the title is in one of the chunks
				titleFound := false

				for _, chunk := range chunks {
					if strings.Contains(chunk, "# Document Title With Space") {
						titleFound = true
						break
					}
				}

				if !titleFound {
					t.Error("Expected one of the chunks to contain the title from front matter")
				}
			},
		},
		{
			name: "Large markdown that might be split into multiple chunks",
			content: strings.Repeat(`
# Section
This is a section with some content.

## Subsection
This is a subsection with more content.

`, 100), // Repeat to create a large document
			expectedChunks: 100 * 2,
			checkFunc: func(t *testing.T, chunks []string) {
				if len(chunks) == 0 {
					t.Fatal("Expected at least one chunk")
				}
				// Check that each chunk contains some expected content
				for i, chunk := range chunks {
					if !strings.Contains(chunk, "# Section") && !strings.Contains(chunk, "## Subsection") {
						t.Errorf("Chunk %d does not contain expected content", i)
					}
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function
			chunks, err := splitMarkdown(tc.content)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			// Check the number of chunks if expected
			if tc.expectedChunks > 0 && len(chunks) != tc.expectedChunks {
				t.Errorf("Expected %d chunks, got: %d", tc.expectedChunks, len(chunks))
			}

			// Run additional checks
			if tc.checkFunc != nil {
				tc.checkFunc(t, chunks)
			}
		})
	}
}
