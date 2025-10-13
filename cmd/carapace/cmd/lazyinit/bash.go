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

  local command="${COMP_WORDS[0]}" nospace data compline="${COMP_LINE:0:${COMP_POINT}}"

  if echo ${compline}"''" | xargs echo 2>/dev/null > /dev/null; then
  	data=$(echo ${compline}"''" | xargs carapace "${command}" bash)
  elif echo ${compline} | sed "s/\$/'/" | xargs echo 2>/dev/null > /dev/null; then
  	data=$(echo ${compline} | sed "s/\$/'/" | xargs carapace "${command}" bash)
  else
  	data=$(echo ${compline} | sed 's/$/"/' | xargs carapace "${command}" bash)
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
