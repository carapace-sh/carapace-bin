package lazyinit

import (
	"fmt"
	"strings"
)

func Zsh(completers []string) string {
	for i, completer := range completers {
		completers[i] = fmt.Sprintf("%q", completer)
	}
	return fmt.Sprintf(`%v%v

function _carapace_completer {
  local command="$(basename $words[1])"
  local compline=${words[@]:0:$CURRENT}
  local IFS=$'\n'
  local lines

  declare -x CARAPACE_COMPLINE="${words}"
  declare -x CARAPACE_ZSH_HASH_DIRS="$(hash -d)"
  declare -x CARAPACE_SHELL=zsh
  declare -x CARAPACE_SHELL_BUILTINS="$(print -roC1 -- ${(k)builtins})"
  declare -x CARAPACE_SHELL_FUNCTIONS="$(print -l ${(ok)functions})"

  # shellcheck disable=SC2086,SC2154,SC2155
  lines="$(echo "${compline}''" | xargs carapace "${command}" zsh 2>/dev/null)"
  if [ $? -eq 1 ]; then
    lines="$(echo "${compline}'" | xargs carapace "${command}" zsh 2>/dev/null)"
    if [ $? -eq 1 ]; then
      lines="$(echo "${compline}\"" | xargs carapace "${command}" zsh 2>/dev/null)"
    fi
  fi

  local zstyle message data
  IFS=$'\001' read -r -d '' zstyle message data <<<"${lines}"
  # shellcheck disable=SC2154
  zstyle ":completion:${curcontext}:*" list-colors "${zstyle}"
  zstyle ":completion:${curcontext}:*" group-name ''
  [ -z "$message" ] || _message -r "${message}"

  local block tag displays values displaysArr valuesArr
  while IFS=$'\002' read -r -d $'\002' block; do
    IFS=$'\003' read -r -d '' tag displays values <<<"${block}"
    # shellcheck disable=SC2034
    IFS=$'\n' read -r -d $'\004' -A displaysArr <<<"${displays}"$'\004'
    IFS=$'\n' read -r -d $'\004' -A valuesArr <<<"${values}"$'\004'

    [[ ${#valuesArr[@]} -gt 1 ]] && _describe -t "${tag}" "${tag}" displaysArr valuesArr -Q -S ''
  done <<<"${data}"
}

compdef _carapace_completer %v
`, pathSnippet("zsh"), envSnippet("zsh"), strings.Join(completers, " "))
}
