package snippet

import (
	"strings"
	"testing"

	shlex "github.com/carapace-sh/carapace-shlex/v2"
)

func TestTcshCompleterScreening(t *testing.T) {
	benign := "my.tool-2"
	hostile := []string{"a`touch X`b", "b$(x)c", "c'; d", "e\"f", "sp ace", "$var"}

	out := Tcsh(append([]string{benign}, hostile...))

	if strings.Count(out, "complete") != 1 {
		t.Fatalf("expected exactly one complete line: %q", out)
	}

	full, err := shlex.Split(out, shlex.Tcsh)
	if err != nil {
		t.Fatalf("tcsh Split failed: %v", err)
	}

	got := make([]string, 0, len(full))
	for _, token := range full {
		if token.Value == "" {
			continue
		}
		got = append(got, token.Value)
	}
	want := []string{"complete", benign,
		"p@*@`echo \"$COMMAND_LINE''\" | xargs env CARAPACE_SHELL=tcsh carapace my.tool-2 tcsh `@@", ";"}
	if len(got) != len(want) {
		t.Fatalf("token count = %v, want %v: %v", len(got), len(want), full)
	}
	for index, value := range got {
		if value != want[index] {
			t.Fatalf("token %v = %q, want %q", index, value, want[index])
		}
	}
}
