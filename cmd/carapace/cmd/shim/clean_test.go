package shim

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestRemoveObsolete(t *testing.T) {
	configDir := t.TempDir()
	binDir := filepath.Join(configDir, "carapace", "bin")
	specDir := filepath.Join(configDir, "carapace", "specs")
	for _, dir := range []string{binDir, specDir} {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatal(err)
		}
	}

	extension := ""
	if runtime.GOOS == "windows" {
		extension = ".exe"
	}
	staleShim := filepath.Join(binDir, "stale"+extension)
	currentShim := filepath.Join(binDir, "current"+extension)
	for _, path := range []string{staleShim, currentShim} {
		if err := os.WriteFile(path, nil, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(specDir, "current.yaml"), nil, 0o644); err != nil {
		t.Fatal(err)
	}

	if err := removeObsoleteIn(configDir); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(staleShim); !os.IsNotExist(err) {
		t.Fatalf("obsolete shim still exists: %v", err)
	}
	if _, err := os.Stat(currentShim); err != nil {
		t.Fatalf("current shim was removed: %v", err)
	}
}
