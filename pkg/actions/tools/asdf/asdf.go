// package asdf contains asdf related actions
package asdf

import (
	"os"
	"path/filepath"
	"strings"
)

func dataDir() string {
	dir := os.Getenv("ASDF_DATA_DIR")
	if dir == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		return filepath.Join(home, ".asdf")
	}

	if dir == "~" || strings.HasPrefix(dir, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return ""
		}
		dir = filepath.Join(home, dir[1:])
	}
	return filepath.Clean(dir)
}
