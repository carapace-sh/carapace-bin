package util

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func FindReverse(path string, name string) (target string, err error) {
	var absPath string
	if absPath, err = filepath.Abs(path); err == nil {
		target = absPath + "/" + name
		if _, err = os.Stat(target); err != nil {
			parent := filepath.Dir(absPath)
			if parent != path {
				return FindReverse(parent, name)
			} else {
				err = errors.New("could not find: " + name)
			}
		}
	}
	return
}

func HasPathPrefix(s string) bool {
	return strings.HasPrefix(s, ".") ||
		strings.HasPrefix(s, "/") ||
		strings.HasPrefix(s, "~")
}
