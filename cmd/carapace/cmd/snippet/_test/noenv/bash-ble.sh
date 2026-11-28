[[ ${BASH_VERSION:-} == 3.* ]] && { printf 'Error: legacy bash is not supported\n' >&2; return 1; }
# export PATH="/tmp/carapace-snippet-fixtures/carapace/bin:$PATH"

_carapace_completer() {
  export COMP_LINE
  export COMP_POINT
  export COMP_TYPE
  export COMP_WORDBREAKS

  declare -x CARAPACE_SHELL=bash

  local command="${COMP_WORDS[0]}" nospace data compline="${COMP_LINE:0:${COMP_POINT}}"

  data=$(echo "${compline}''" | xargs carapace "${command}" bash 2>/dev/null)
  local _carapace_rc=$?
  if [ $_carapace_rc -eq 1 ]; then
    data=$(echo "${compline}'" | xargs carapace "${command}" bash 2>/dev/null)
    _carapace_rc=$?
    if [ $_carapace_rc -eq 1 ]; then
    	data=$(echo "${compline}\"" | xargs carapace "${command}" bash 2>/dev/null)
    	_carapace_rc=$?
    fi
  fi
  if [ $_carapace_rc -ne 0 ] && [ -n "${CARAPACE_SHELL_STATE:-}" ]; then
    local _carapace_old_state="${CARAPACE_SHELL_STATE}"
    unset CARAPACE_SHELL_STATE
    data=$(echo "${compline}''" | xargs carapace "${command}" bash 2>/dev/null)
    _carapace_rc=$?
    if [ $_carapace_rc -eq 1 ]; then
      data=$(echo "${compline}'" | xargs carapace "${command}" bash 2>/dev/null)
      _carapace_rc=$?
      if [ $_carapace_rc -eq 1 ]; then
      	data=$(echo "${compline}\"" | xargs carapace "${command}" bash 2>/dev/null)
      fi
    fi
    declare -x CARAPACE_SHELL_STATE="${_carapace_old_state}"
  fi

  IFS=$'\001' read -r -d '' nospace data <<<"${data}"
  mapfile -t COMPREPLY < <(echo "${data}")
  unset COMPREPLY[-1]

  [ "${nospace}" = true ] && compopt -o nospace
  local IFS=$'\n'
  [[ "${COMPREPLY[*]}" == "" ]] && COMPREPLY=() # fix for mapfile creating a non-empty array from empty command output
  [ -n "${_carapace_state_cleanup:-}" ] && rm -rf "$_carapace_state_cleanup"
}

_carapace_completer_ble() {
  if [[ ${BLE_ATTACHED-} ]]; then
    [[ :$comp_type: == *:auto:* ]] && return

    compopt -o ble/no-default
    bleopt complete_menu_style=desc

	local command="${COMP_WORDS[0]}"
    local compline="${COMP_LINE:0:${COMP_POINT}}"
    local IFS=$'\n'
    local c
    mapfile -t c < <(echo "$compline" | sed -e "s/ \$/ ''/" -e 's/"/\"/g' | xargs carapace "${command}" bash-ble 2>/dev/null)
    local _carapace_rc=$?
    if [ $_carapace_rc -ne 0 ] && [ -n "${CARAPACE_SHELL_STATE:-}" ]; then
      local _carapace_old_state="${CARAPACE_SHELL_STATE}"
      unset CARAPACE_SHELL_STATE
      mapfile -t c < <(echo "$compline" | sed -e "s/ \$/ ''/" -e 's/"/\"/g' | xargs carapace "${command}" bash-ble 2>/dev/null)
      declare -x CARAPACE_SHELL_STATE="${_carapace_old_state}"
    fi
    [[ "${c[*]}" == "" ]] && c=() # fix for mapfile creating a non-empty array from empty command output

    local cand
    for cand in "${c[@]}"; do
      [ ! -z "$cand" ] && ble/complete/cand/yield mandb "${cand%$'\t'*}" "${cand##*$'\t'}"
    done
  else
    _carapace_completer
  fi
}

complete -F _carapace_completer_ble example

