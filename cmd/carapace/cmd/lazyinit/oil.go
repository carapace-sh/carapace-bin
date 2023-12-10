package lazyinit

import (
	"fmt"
	"strings"
)

func Oil(completers []string) string {
	snippet := `%v%v

_carapace_lazy() {
  source <(carapace $1 oil)
   $"_$1_completion"
}
complete -F _carapace_lazy %v
`
	return fmt.Sprintf(snippet, pathSnippet("oil"), envSnippet("oil"), strings.Join(completers, " "))
}
