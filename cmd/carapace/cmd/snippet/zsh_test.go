package snippet

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestZshCommandSubstitutionWords(t *testing.T) {
	if _, err := exec.LookPath("zsh"); err != nil {
		t.Skip("zsh not available")
	}

	tests := map[string]string{
		"dollar_parens": "$(git rev-parse --show-toplevel)",
		"backticks":     "`git rev-parse --show-toplevel`",
	}

	for name, substitution := range tests {
		t.Run(name, func(t *testing.T) {
			dir := t.TempDir()
			argsPath := filepath.Join(dir, "args")
			scriptPath := filepath.Join(dir, "complete.zsh")

			script := strings.Join([]string{
				`compdef() { :; }`,
				`zstyle() { :; }`,
				`_message() { :; }`,
				`_describe() { :; }`,
				`carapace() { print -r -l -- "$@" >| "$CARAPACE_TEST_ARGS"; }`,
				Zsh([]string{"hx"}),
				`words=(hx --working-dir '` + substitution + `' --)`,
				`CURRENT=4`,
				`curcontext=carapace-test`,
				`_carapace_completer`,
			}, "\n")

			if err := os.WriteFile(scriptPath, []byte(script), 0o600); err != nil {
				t.Fatal(err)
			}

			cmd := exec.Command("zsh", "-f", scriptPath)
			cmd.Env = append(os.Environ(), "CARAPACE_TEST_ARGS="+argsPath)
			if output, err := cmd.CombinedOutput(); err != nil {
				t.Fatalf("zsh completion failed: %v\n%s", err, output)
			}

			content, err := os.ReadFile(argsPath)
			if err != nil {
				t.Fatal(err)
			}
			got := strings.TrimSuffix(string(content), "\n")
			want := strings.Join([]string{
				"hx",
				"zsh",
				"hx",
				"--working-dir",
				substitution,
				"--",
			}, "\n")
			if got != want {
				t.Fatalf("args mismatch:\nwant:\n%s\n\ngot:\n%s", want, got)
			}
		})
	}
}
