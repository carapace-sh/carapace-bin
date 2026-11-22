package snippet

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func testSnippetFixture(t *testing.T, shell string, file string, content string) {
	t.Helper()
	b, err := os.ReadFile(file)
	if err != nil {
		t.Fatal("failed to read fixture file")
	}
	expected := strings.ReplaceAll(string(b), "\r\n", "\n")
	if configHome := os.Getenv("XDG_CONFIG_HOME"); configHome != "" {
		expected = strings.ReplaceAll(expected, "/tmp/carapace-snippet-fixtures", filepath.ToSlash(configHome))
	}
	if runtime.GOOS == "windows" && expected != content+"\n" {
		t.Skipf("fixture does not cover windows-specific snippet: %v", shell)
	}
	if expected != content+"\n" {
		t.Fatalf("snippet mismatch for %v", shell)
	}
}

func snippetFixtures(t *testing.T, env string, builtins func(shell string) string) map[string]string {
	t.Helper()
	t.Setenv("CARAPACE_ENV", env)

	configDir := filepath.Join(os.TempDir(), "carapace-snippet-fixtures")
	binDir := filepath.Join(configDir, "carapace", "bin")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("XDG_CONFIG_HOME", configDir)
	t.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH"))

	completers := []string{"example"}
	fixtures := map[string]string{}
	for _, shell := range []string{"bash", "bash-ble", "cmd-clink", "elvish", "fish", "nushell", "oil", "powershell", "tcsh", "xonsh", "zsh"} {
		t.Setenv("CARAPACE_BUILTINS", builtins(shell))
		switch shell {
		case "bash":
			fixtures[shell] = Bash(completers)
		case "bash-ble":
			fixtures[shell] = BashBle(completers)
		case "cmd-clink":
			fixtures[shell] = CmdClink(completers)
		case "elvish":
			fixtures[shell] = Elvish(completers)
		case "fish":
			fixtures[shell] = Fish(completers)
		case "nushell":
			fixtures[shell] = Nushell(completers)
		case "oil":
			fixtures[shell] = Oil(completers)
		case "powershell":
			fixtures[shell] = Powershell(completers)
		case "tcsh":
			fixtures[shell] = Tcsh(completers)
		case "xonsh":
			fixtures[shell] = Xonsh(completers)
		case "zsh":
			fixtures[shell] = Zsh(completers)
		}
	}
	return fixtures
}

func TestSnippetFixtures(t *testing.T) {
	files := map[string]string{
		"bash":       "bash.sh",
		"bash-ble":   "bash-ble.sh",
		"cmd-clink":  "cmd-clink.lua",
		"elvish":     "elvish.elv",
		"fish":       "fish.fish",
		"nushell":    "nushell.nu",
		"oil":        "oil.sh",
		"powershell": "powershell.ps1",
		"tcsh":       "tcsh.tcsh",
		"xonsh":      "xonsh.py",
		"zsh":        "zsh.sh",
	}
	variants := []struct {
		name     string
		env      string
		builtins func(string) string
	}{
		{"noenv", "0", func(string) string { return "" }},
		{"env", "1", func(string) string { return "" }},
		{"builtins", "0", func(shell string) string { return shell }},
		{"env_builtins", "1", func(shell string) string { return shell }},
	}

	for _, variant := range variants {
		variant := variant
		t.Run(variant.name, func(t *testing.T) {
			fixtures := snippetFixtures(t, variant.env, variant.builtins)
			for shell, file := range files {
				shell, file := shell, file
				t.Run(shell, func(t *testing.T) {
					testSnippetFixture(t, shell, filepath.Join("_test", variant.name, file), fixtures[shell])
				})
			}
		})
	}
}
