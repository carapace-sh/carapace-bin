package snippet

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	shlex "github.com/carapace-sh/carapace-shlex/v2"
)

func TestPathSnippetEscaping(t *testing.T) {
	configHome := "/opt/\"d$(id)'" + "`" + "pa/th\\\\x\t\nmore"
	if runtime.GOOS == "windows" {
		// must be an absolute path: xdg.UserConfigDir falls back to the OS
		// default directory otherwise (and forwards slashes paths anyway)
		configHome = "C:" + configHome
	}
	t.Setenv("XDG_CONFIG_HOME", configHome)

	binDir := filepath.ToSlash(configHome) + "/carapace/bin"

	for _, shell := range []string{"bash", "bash-ble", "oil", "zsh", "elvish", "fish", "nushell", "powershell", "xonsh"} {
		t.Run(shell, func(t *testing.T) {
			snippet := pathSnippet(shell)
			if !strings.Contains(snippet, "carapace") {
				t.Fatalf("unexpected snippet: %q", snippet)
			}

			var format shlex.Format
			switch shell {
			case "bash", "bash-ble":
				format = shlex.Bash
			case "oil":
				format = shlex.Oil
			case "zsh":
				format = shlex.Zsh
			case "elvish":
				format = shlex.Elvish
			case "fish":
				format = shlex.Fish
			case "nushell":
				format = shlex.Nushell
			case "powershell":
				format = shlex.Powershell
			case "xonsh":
				format = shlex.Xonsh
			}

			tokens, err := shlex.Split(snippet, format)
			if err != nil {
				t.Fatalf("Split failed: %v", err)
			}

			expected := binDir
			if shell == "xonsh" {
				// shlex's xonsh lexer keeps the \n and \r escapes literally
				// (Value differs from the actual xonsh value)
				expected = strings.NewReplacer("\n", `\n`, "\r", `\r`).Replace(binDir)
			}

			found := false
			for _, token := range tokens {
				if strings.Contains(token.Value, expected) {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("no token containing literal binDir %q after splitting : %q", expected, snippet)
			}
		})
	}
}
