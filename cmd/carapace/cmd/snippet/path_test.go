package snippet

import (
	"strings"
	"testing"

	shlex "github.com/carapace-sh/carapace-shlex/v2"
)

func TestPathSnippetEscaping(t *testing.T) {
	configHome := "/opt/\"d$(id)'" + "`" + "pa/th\\\\x"
	t.Setenv("XDG_CONFIG_HOME", configHome)

	binDir := configHome + "/carapace/bin"

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

			found := false
			for _, token := range tokens {
				if strings.Contains(token.Value, binDir) {
					found = true
					break
				}
			}
			if !found {
				t.Fatalf("no token containing literal binDir %q after splitting : %q", binDir, snippet)
			}
		})
	}
}
