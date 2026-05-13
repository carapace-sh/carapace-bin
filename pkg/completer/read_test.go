package completer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadSpec(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cmd.yaml")
	if err := os.WriteFile(path, []byte("name: cmd\ndescription: test command\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	specs, err := ReadSpec(path, "user", true)
	if err != nil {
		t.Fatal(err)
	}

	matches := specs["cmd"]
	if len(matches) != 1 {
		t.Fatalf("expected one spec, got %d", len(matches))
	}
	if matches[0].Description != "test command" {
		t.Fatalf("unexpected description: %q", matches[0].Description)
	}
	if matches[0].Spec != path {
		t.Fatalf("unexpected spec path: %q", matches[0].Spec)
	}
}

func TestReadSpecFilenameMustMatchName(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "cmd.yaml")
	if err := os.WriteFile(path, []byte("name: other\n"), 0o644); err != nil {
		t.Fatal(err)
	}

	if _, err := ReadSpec(path, "user", true); err == nil {
		t.Fatal("expected filename/name mismatch error")
	}
}
