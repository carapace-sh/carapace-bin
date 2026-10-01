package snippet

import (
	"os/exec"
	"strings"
	"testing"
)

func TestZshQuotesWordsBeforeXargs(t *testing.T) {
	snippet := Zsh([]string{"hx"})

	for _, expected := range []string{
		"local raw_compline=${words[@]:0:$CURRENT}",
		"local compline=${(j: :)${(qq)${words[@]:0:$CURRENT}}}",
		`declare -x CARAPACE_COMPLINE="${raw_compline}"`,
	} {
		if !strings.Contains(snippet, expected) {
			t.Fatalf("zsh snippet missing %q", expected)
		}
	}
}

func TestZshQuotedComplineKeepsInterpolationArgument(t *testing.T) {
	if _, err := exec.LookPath("zsh"); err != nil {
		t.Skip("zsh not available")
	}

	output, err := exec.Command("zsh", "-fc", `
words=(hx --working-dir "\$(git rev-parse --show-toplevel)" --)
CURRENT=4
compline=${(j: :)${(qq)${words[@]:0:$CURRENT}}}
print -r -- "${compline}''" | xargs -n1 printf '<%s>\n'
`).CombinedOutput()
	if err != nil {
		t.Fatalf("zsh command failed: %v\n%s", err, output)
	}

	expected := "<hx>\n<--working-dir>\n<$(git rev-parse --show-toplevel)>\n<-->\n"
	if string(output) != expected {
		t.Fatalf("unexpected xargs output\nactual:\n%sexpected:\n%s", output, expected)
	}
}
