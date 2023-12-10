package lazyinit

import (
	"fmt"
	"strings"
)

func Tcsh(completers []string) string {
	// TODO hardcoded for now
	snippet := make([]string, len(completers))
	for index, c := range completers {
		snippet[index] = fmt.Sprintf("complete \"%v\" 'p@*@`echo \"$COMMAND_LINE'\"''\"'\" | xargs carapace %v tcsh `@@' ;", c, c)
	}
	return strings.Join(snippet, "\n")
}
