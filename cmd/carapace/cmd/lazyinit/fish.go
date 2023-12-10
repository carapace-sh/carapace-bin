package lazyinit

import (
	"fmt"
	"strings"
)

func Fish(completers []string) string {
	snippet := `%v%v

function _carapace_lazy
   complete -c $argv[1] -e
   carapace $argv[1] fish | source
   complete --do-complete=(commandline -cp)
end
%v
`
	complete := make([]string, len(completers))
	for index, completer := range completers {
		complete[index] = fmt.Sprintf(`complete -c '%v' -f -a '(_carapace_lazy %v)'`, completer, completer)
	}
	return fmt.Sprintf(snippet, pathSnippet("fish"), envSnippet("fish"), strings.Join(complete, "\n"))
}
