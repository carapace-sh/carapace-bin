[[ ${BASH_VERSION:-} == 3.* ]] && { printf 'Error: legacy bash is not supported\n' >&2; return 1; }
# export PATH="/tmp/carapace-snippet-fixtures/carapace/bin:$PATH"

_carapace_completer() {
  export COMP_LINE
  export COMP_POINT
  export COMP_TYPE
  export COMP_WORDBREAKS

  declare -x CARAPACE_SHELL=bash
  local _carapace_state _carapace_state_cleanup
  _carapace_state=""
  _carapace_state_cleanup=""
  if [ -n "${XDG_RUNTIME_DIR:-}" ] && [ -d "${XDG_RUNTIME_DIR}" ] && [ -O "${XDG_RUNTIME_DIR}" ] && [ ! -L "${XDG_RUNTIME_DIR}" ]; then
    _carapace_state="${XDG_RUNTIME_DIR}/carapace/builtins/${UID:-$(id -u)}/$$"
  else
    _carapace_state="$(mktemp -d "${TMPDIR:-/tmp}/carapace-builtin.XXXXXX" 2>/dev/null)" || _carapace_state=""
    _carapace_state_cleanup="$_carapace_state"
  fi
  if [ -n "$_carapace_state" ] && mkdir -m 700 -p "$_carapace_state" 2>/dev/null && [ -O "$_carapace_state" ] && [ ! -L "$_carapace_state" ]; then
    declare -x CARAPACE_SHELL_STATE="$_carapace_state"
    compgen -a > "$_carapace_state/aliases"
    compgen -b > "$_carapace_state/builtins"
    compgen -A function > "$_carapace_state/functions"
    jobs 2>/dev/null | while read -r line; do [[ $line =~ \[([0-9]+)\] ]] && echo "%${BASH_REMATCH[1]}"; done > "$_carapace_state/jobs"
    compgen -v > "$_carapace_state/variables"
  else
    unset CARAPACE_SHELL_STATE
  fi

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

complete -o noquote -F _carapace_completer example

