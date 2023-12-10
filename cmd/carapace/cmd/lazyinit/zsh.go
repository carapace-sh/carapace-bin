package lazyinit

import (
	"fmt"
	"strings"
)

func Zsh(completers []string) string {
	snippet := `%v%v

function _carapace_lazy {
    source <(carapace $words[1] zsh)
}
compdef _carapace_lazy %v
`
	return fmt.Sprintf(snippet, pathSnippet("zsh"), envSnippet("zsh"), strings.Join(completers, " "))
}
