package resources

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/tmc/langchaingo/textsplitter"
)

const (
	maxMarkdownChunkSize    = 65536
	maxMarkdownChunkOverlap = 8192
)

func glob(fsDir fs.FS, rootPath string, fn func(string) bool) ([]string, error) {
	var files []string
	if err := fs.WalkDir(fsDir, rootPath, func(s string, _ fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if fn(s) {
			files = append(files, s)
		}
		return nil
	}); err != nil {
		return nil, fmt.Errorf("error walking directory: %w", err)
	}
	return files, nil
}

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
