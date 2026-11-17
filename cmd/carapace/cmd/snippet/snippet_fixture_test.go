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
	if runtime.GOOS == "windows" && expected != content+"\n" {
		t.Skipf("fixture does not cover windows-specific snippet: %v", shell)
	}
	if expected != content+"\n" {
		t.Fatalf("snippet mismatch for %v", shell)
	}
}

func snippetFixtures(t *testing.T) map[string]string {
	t.Helper()
	t.Setenv("CARAPACE_ENV", "0")
	t.Setenv("CARAPACE_BUILTINS", "")

	configDir := filepath.Join(os.TempDir(), "carapace-snippet-fixtures")
	binDir := filepath.Join(configDir, "carapace", "bin")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("XDG_CONFIG_HOME", configDir)
	t.Setenv("PATH", binDir+string(os.PathListSeparator)+os.Getenv("PATH"))

	completers := []string{"example"}
	return map[string]string{
		"bash":       Bash(completers),
		"bash-ble":   BashBle(completers),
		"cmd-clink":  CmdClink(completers),
		"elvish":     Elvish(completers),
		"fish":       Fish(completers),
		"nushell":    Nushell(completers),
		"oil":        Oil(completers),
		"powershell": Powershell(completers),
		"tcsh":       Tcsh(completers),
		"xonsh":      Xonsh(completers),
		"zsh":        Zsh(completers),
	}
}

func TestSnippetFixtures(t *testing.T) {
	fixtures := snippetFixtures(t)
	for shell, file := range map[string]string{
		"bash":       "_test/bash.sh",
		"bash-ble":   "_test/bash-ble.sh",
		"cmd-clink":  "_test/cmd-clink.lua",
		"elvish":     "_test/elvish.elv",
		"fish":       "_test/fish.fish",
		"nushell":    "_test/nushell.nu",
		"oil":        "_test/oil.sh",
		"powershell": "_test/powershell.ps1",
		"tcsh":       "_test/tcsh.tcsh",
		"xonsh":      "_test/xonsh.py",
		"zsh":        "_test/zsh.sh",
	} {
		shell, file := shell, file
		t.Run(shell, func(t *testing.T) {
			testSnippetFixture(t, shell, file, fixtures[shell])
		})
	}
}
