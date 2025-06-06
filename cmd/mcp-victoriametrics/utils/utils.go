package utils

import (
	"fmt"
	"io/fs"
)

func Glob(fsDir fs.FS, rootPath string, fn func(string) bool) ([]string, error) {
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
