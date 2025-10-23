package lazyinit

import (
	"fmt"
	"strings"
)

const bashCompleter = `_carapace_completer() {
  export COMP_LINE
  export COMP_POINT
  export COMP_TYPE
  export COMP_WORDBREAKS

  declare -x CARAPACE_SHELL=bash
  declare -x CARAPACE_SHELL_BUILTINS="$(compgen -b)"
  declare -x CARAPACE_SHELL_FUNCTIONS="$(compgen -A function)"

  local command="${COMP_WORDS[0]}" nospace data compline="${COMP_LINE:0:${COMP_POINT}}"

  data=$(echo "${compline}''" | xargs carapace "${command}" bash 2>/dev/null)
  if [ $? -eq 1 ]; then
    data=$(echo "${compline}'" | xargs carapace "${command}" bash 2>/dev/null)
    if [ $? -eq 1 ]; then
    	data=$(echo "${compline}\"" | xargs carapace "${command}" bash 2>/dev/null)
    fi
  fi

  IFS=$'\001' read -r -d '' nospace data <<<"${data}"
  mapfile -t COMPREPLY < <(echo "${data}")
  unset COMPREPLY[-1]

  [ "${nospace}" = true ] && compopt -o nospace
  local IFS=$'\n'
  [[ "${COMPREPLY[*]}" == "" ]] && COMPREPLY=() # fix for mapfile creating a non-empty array from empty command output
}
`

func Bash(completers []string) string {
	for i, completer := range completers {
		completers[i] = fmt.Sprintf("%q", completer)
	}
	return fmt.Sprintf(`%v%v

%v
complete -o noquote -F _carapace_completer %v
`, pathSnippet("bash"), envSnippet("bash"), bashCompleter, strings.Join(completers, " "))
}
