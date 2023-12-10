package lazyinit

import (
	"fmt"
	"strings"
)

func Bash(completers []string) string {
	snippet := `%v%v

_carapace_lazy() {
  source <(carapace $1 bash)
   $"_$1_completion"
}
complete -F _carapace_lazy %v
`
	return fmt.Sprintf(snippet, pathSnippet("bash"), envSnippet("bash"), strings.Join(completers, " "))
}

func BashBle(completers []string) string {
	snippet := `%v

_carapace_lazy() {
  source <(carapace $1 bash-ble)
   $"_$1_completion_ble"
}
complete -F _carapace_lazy %v
`
	return fmt.Sprintf(snippet, pathSnippet("bash-ble"), strings.Join(completers, " "))
}
