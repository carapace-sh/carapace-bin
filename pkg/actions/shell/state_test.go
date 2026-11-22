package shell

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestReadShellKeyStateDir(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "functions"), []byte("foo\n_bar\n\nbaz\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv(shellStateEnv, dir)
	t.Setenv("CARAPACE_SHELL_FUNCTIONS", "fallback")

	got := readShellKey("functions")
	want := []string{"foo", "_bar", "baz"}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestReadShellKeyEnvFallback(t *testing.T) {
	t.Setenv(shellStateEnv, "")
	t.Setenv("CARAPACE_SHELL_BUILTINS", "echo cd")

	got := readShellKey("builtins")
	want := []string{"echo", "cd"}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestReadShellKeyInvalidStateDir(t *testing.T) {
	t.Setenv(shellStateEnv, "relative/path")
	t.Setenv("CARAPACE_SHELL_ALIASES", "ll")

	got := readShellKey("aliases")
	want := []string{"ll"}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestReadShellKeyMissingFileFallsBack(t *testing.T) {
	dir := t.TempDir()
	t.Setenv(shellStateEnv, dir)
	t.Setenv("CARAPACE_SHELL_JOBS", "%1 %2")

	got := readShellKey("jobs")
	want := []string{"%1", "%2"}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestReadShellKeyRejectsSymlink(t *testing.T) {
	dir := t.TempDir()
	target := filepath.Join(dir, "target")
	if err := os.WriteFile(target, []byte("secret\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(target, filepath.Join(dir, "functions")); err != nil {
		t.Skip(err)
	}
	t.Setenv(shellStateEnv, dir)
	t.Setenv("CARAPACE_SHELL_FUNCTIONS", "fallback")

	got := readShellKey("functions")
	want := []string{"fallback"}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("got %v want %v", got, want)
	}
}

func TestReadShellKeySizeLimit(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "functions"), []byte(strings.Repeat("a\n", maxStateLines+10)), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv(shellStateEnv, dir)

	got := readShellKey("functions")
	if len(got) != maxStateLines {
		t.Fatalf("got %d want %d", len(got), maxStateLines)
	}
}
