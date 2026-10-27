package snippet

import (
	"fmt"
	"strings"

	shlex "github.com/carapace-sh/carapace-shlex/v2"
)

func Tcsh(completers []string) string {
	// TODO hardcoded for now
	snippet := make([]string, 0, len(completers))
	for _, c := range completers {
		// the completer name is interpolated twice into a tcsh complete line whose
		// second occurrence sits inside a hand-built quote sequence that cannot
		// represent arbitrary characters, so only inert barewords are supported
		if !shlex.IsBareWord(c, shlex.Tcsh) {
			continue
		}
		snippet = append(snippet, fmt.Sprintf("complete \"%v\" 'p@*@`echo \"$COMMAND_LINE'\"''\"'\" | xargs env CARAPACE_SHELL=tcsh carapace %v tcsh `@@' ;", c, c))
	}
	return strings.Join(snippet, "\n")
}
