package lazyinit

import (
	"fmt"
	"strings"
)

func BashBle(completers []string) string {
	for i, completer := range completers {
		completers[i] = fmt.Sprintf("%q", completer)
	}
	return fmt.Sprintf(`%v

%v
_carapace_completer_ble() {
  if [[ ${BLE_ATTACHED-} ]]; then
    [[ :$comp_type: == *:auto:* ]] && return

    compopt -o ble/no-default
    bleopt complete_menu_style=desc

	local command="${COMP_WORDS[0]}"
    local compline="${COMP_LINE:0:${COMP_POINT}}"
    local IFS=$'\n'
    local c
    mapfile -t c < <(echo "$compline" | sed -e "s/ \$/ ''/" -e 's/"/\"/g' | xargs carapace "${command}" bash-ble)
    [[ "${c[*]}" == "" ]] && c=() # fix for mapfile creating a non-empty array from empty command output

    local cand
    for cand in "${c[@]}"; do
      [ ! -z "$cand" ] && ble/complete/cand/yield mandb "${cand%%$'\t'*}" "${cand##*$'\t'}"
    done
  else
    _carapace_completer
  fi
}

complete -F _carapace_completer_ble %v
`, pathSnippet("bash-ble"), bashCompleter, strings.Join(completers, " "))
}
