# export PATH="/tmp/carapace-snippet-fixtures/carapace/bin:$PATH"

_carapace_completer() {
  declare -x CARAPACE_SHELL=oil

  local command="${COMP_WORDS[0]}"
  local compline="${COMP_LINE:0:${COMP_POINT}}"
  local IFS=$'\n'
  mapfile -t COMPREPLY < <(echo "$compline" | sed -e "s/ \$/ ''/" -e 's/"/\"/g' | xargs carapace "${command}" oil)
  [[ "${COMPREPLY[@]}" == "" ]] && COMPREPLY=() # fix for mapfile creating a non-empty array from empty command output
  [[ ${COMPREPLY[0]} == *[/=@:.,$'\001'] ]] && compopt -o nospace
  # TODO use mapfile
  # shellcheck disable=SC2206
  [[ ${#COMPREPLY[@]} -eq 1 ]] && COMPREPLY=(${COMPREPLY[@]%$'\001'})
}

complete -F _carapace_completer example

