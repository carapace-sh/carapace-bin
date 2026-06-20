package snippet

import (
	"fmt"
	"strings"
)

func Tcsh(completers []string) string {
	snippet := make([]string, len(completers))
	for index, c := range completers {
		snippet[index] = fmt.Sprintf("complete \"%v\" 'p@*@`echo \\\"${COMMAND_LINE}\\\"''\\\"'\\\" | xargs env CARAPACE_SHELL=tcsh CARAPACE_SHELL_BUILTINS=\\\"$(compgen -b | tr '\\\\n' ' ')\\\" CARAPACE_SHELL_FUNCTIONS=\\\"$(compgen -A function | tr '\\\\n' ' ')\\\" CARAPACE_SHELL_VARIABLES=\\\"$(compgen -v | tr '\\\\n' ' ')\\\" carapace %v tcsh `@@' ;", c, c)
	}
	return strings.Join(snippet, "\n")
}
