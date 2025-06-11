package resources

import (
	"fmt"
	"strings"

	"github.com/tmc/langchaingo/textsplitter"
)

const (
	maxMarkdownChunkSize    = 65536
	maxMarkdownChunkOverlap = 8192
)

var (
	mdSplitter = textsplitter.NewMarkdownTextSplitter(
		textsplitter.WithCodeBlocks(true),
		textsplitter.WithHeadingHierarchy(true),
		textsplitter.WithJoinTableRows(true),
		textsplitter.WithKeepSeparator(true),
		textsplitter.WithReferenceLinks(true),
		textsplitter.WithAllowedSpecial([]string{"all"}),
		textsplitter.WithChunkSize(maxMarkdownChunkSize),
		textsplitter.WithChunkOverlap(maxMarkdownChunkOverlap),
	)
)

func splitMarkdown(content string) ([]string, error) {
	var (
		frontMatter      string
		frontMatterTitle string
	)

	for line := range strings.Lines(content) {
		if len(frontMatter) == 0 {
			if strings.HasPrefix(line, "---") {
				frontMatter += line
				continue
			}
			break
		}
		frontMatter += line
		if strings.HasPrefix(line, "title:") {
			frontMatterTitle = strings.Trim(strings.TrimSpace(strings.TrimPrefix(line, "title:")), "\"'")
		} else if strings.HasPrefix(line, "title :") {
			frontMatterTitle = strings.Trim(strings.TrimSpace(strings.TrimPrefix(line, "title :")), "\"'")
		}
		if strings.HasPrefix(line, "---") {
			break
		}
	}

	content = content[len(frontMatter):]
	if frontMatterTitle != "" {
		content = fmt.Sprintf("# %s\n%s\n", frontMatterTitle, content)
	}

	chunks, err := mdSplitter.SplitText(content)
	if err != nil {
		return nil, fmt.Errorf("error splitting text: %w", err)
	}
	return chunks, nil
}
