package snippet

import (
	"fmt"
	"strings"
)

func Oil(completers []string) string {
	for i, completer := range completers {
		completers[i] = fmt.Sprintf("%q", completer)
	}
	shellVars := `
  export CARAPACE_SHELL=oil
  export CARAPACE_SHELL_ALIASES="$(alias | sed 's/alias //' | sed 's/=.*//')"
  export CARAPACE_SHELL_BUILTINS="$(compgen -b)"
  export CARAPACE_SHELL_FUNCTIONS="$(compgen -A function)"
  export CARAPACE_SHELL_JOBS="$(jobs 2>/dev/null | while read -r line; do [[ $line =~ \[([0-9]+)\] ]] && echo "%${BASH_REMATCH[1]}"; done)"
  export CARAPACE_SHELL_VARIABLES="$(compgen -v)"`

	completerFunc := strings.Replace(`_carapace_completer() {
%v  local command="${COMP_WORDS[0]}"
  local compline="${COMP_LINE:0:${COMP_POINT}}"
  local IFS=$'\n'
  mapfile -t COMPREPLY < <(echo "$compline" | sed -e "s/ \$/ ''/" -e 's/"/\"/g' | xargs carapace "${command}" oil)
  [[ "${COMPREPLY[@]}" == "" ]] && COMPREPLY=() # fix for mapfile creating a non-empty array from empty command output
  [[ ${COMPREPLY[0]} == *[/=@:.,$'\\x01'] ]] && compopt -o nospace
  # TODO use mapfile
  # shellcheck disable=SC2206
  [[ ${#COMPREPLY[@]} -eq 1 ]] && COMPREPLY=(${COMPREPLY[@]%%$'\x01'})
}

complete -F _carapace_completer %v`, "%v", shellVars, 1)

	return fmt.Sprintf("%v%v\n\n%v %v", pathSnippet("oil"), envSnippet("oil"), completerFunc, strings.Join(completers, " "))
}
