package shell

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
)

const (
	shellStateEnv = "CARAPACE_SHELL_STATE"
	maxStateBytes = 1 << 20 // 1 MiB per key
	maxStateLines = 100000
)

func readShellKey(name string) []string {
	if dir := os.Getenv(shellStateEnv); dir != "" {
		if vals, ok := readShellStateKey(dir, name); ok {
			return vals
		}
	}
	return strings.Fields(os.Getenv("CARAPACE_SHELL_" + strings.ToUpper(name)))
}

func readShellStateKey(dir, name string) ([]string, bool) {
	if !filepath.IsAbs(dir) {
		return nil, false
	}

	root, err := os.OpenRoot(filepath.Clean(dir))
	if err != nil {
		return nil, false
	}
	defer root.Close()

	f, err := root.Open(name)
	if err != nil {
		return nil, false
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil || !info.Mode().IsRegular() || info.Size() > maxStateBytes {
		return nil, false
	}

	vals := make([]string, 0, 64)
	scanner := bufio.NewScanner(io.LimitReader(f, maxStateBytes))
	scanner.Buffer(make([]byte, 0, 64*1024), 64*1024)
	for scanner.Scan() {
		if len(vals) >= maxStateLines {
			break
		}
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			vals = append(vals, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, false
	}
	return vals, true
}
