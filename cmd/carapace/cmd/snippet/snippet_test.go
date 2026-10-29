package snippet

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	shlex "github.com/carapace-sh/carapace-shlex/v2"
)

const benignCompleter = "carapace-test-benign"

func TestSnippetScreensCompleterNames(t *testing.T) {
	configDir := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", configDir)
	t.Setenv("XDG_CACHE_HOME", t.TempDir())
	t.Setenv("CARAPACE_BRIDGES", "")
	t.Setenv("CARAPACE_EXCLUDES", "")

	specDir := filepath.Join(configDir, "carapace", "specs")
	if err := os.MkdirAll(specDir, 0o755); err != nil {
		t.Fatal(err)
	}

	hostile := []string{
		"zzz$(touch PROOF)zzz",
		"zzz`touch PROOF`zzz",
		"zzz;zzz",
		"zzz'zzz",
		"zzz zzz",
	}

	for _, name := range append([]string{benignCompleter}, hostile...) {
		body := "name: " + name + "\ndescription: test\n"
		if err := os.WriteFile(filepath.Join(specDir, name+".yaml"), []byte(body), 0o644); err != nil {
			t.Fatalf("write spec %q: %v", name, err)
		}
	}

	for _, shell := range []string{"bash", "bash-ble", "cmd-clink", "elvish", "fish", "oil", "powershell", "tcsh", "xonsh", "zsh"} {
		t.Run(shell, func(t *testing.T) {
			snippet := Snippet(shell)
			if !strings.Contains(snippet, benignCompleter) {
				t.Fatalf("benign completer %q missing from snippet", benignCompleter)
			}
			for _, name := range hostile {
				if strings.Contains(snippet, name) {
					t.Fatalf("completer name %q embedded in snippet", name)
				}
			}
			if shell == "tcsh" {
				assertTcshCompleteLines(t, snippet)
			}
		})
	}
}

func TestTcshCompleteLineQuoting(t *testing.T) {
	benign := "my.tool-2"

	out := Tcsh([]string{benign, "zzz$(touch PROOF)zzz", "zzz`touch PROOF`zzz", "zzz;zzz", "zzz'zzz", "zzz zzz"})

	full, err := shlex.Split(out, shlex.Tcsh)
	if err != nil {
		t.Fatalf("Split(%q): %v", out, err)
	}

	got := make([]string, 0, len(full))
	for _, token := range full {
		if token.Value != "" {
			got = append(got, token.Value)
		}
	}
	want := []string{"complete", benign,
		"p@*@`echo \"$COMMAND_LINE''\" | xargs env CARAPACE_SHELL=tcsh carapace my.tool-2 tcsh `@@", ";"}
	if len(got) != len(want) {
		t.Fatalf("token count = %v, want %v: %v", len(got), len(want), got)
	}
	for index, value := range got {
		if value != want[index] {
			t.Fatalf("token %v = %q, want %q", index, value, want[index])
		}
	}
}

func assertTcshCompleteLines(t *testing.T, snippet string) {
	t.Helper()

	lines := 0
	for _, line := range strings.Split(snippet, "\n") {
		if !strings.HasPrefix(line, "complete ") {
			continue
		}
		lines++
		if _, err := shlex.Split(line, shlex.Tcsh); err != nil {
			t.Fatalf("Split(%q): %v", line, err)
		}
	}
	if lines == 0 {
		t.Fatal("no complete lines generated")
	}
}
