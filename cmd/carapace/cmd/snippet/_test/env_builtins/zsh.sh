# export PATH="/tmp/carapace-snippet-fixtures/carapace/bin:$PATH"

get-env () { echo "${(P)1}"; }
set-env () { export "$1=$2"; }
unset-env () { unset "$1"; }

function _carapace_completer {
  local command="$(basename $words[1])"
  local compline=${words[@]:0:$CURRENT}
  local IFS=$'\n'
  local lines

  declare -x CARAPACE_COMPLINE="${words}"
  declare -x CARAPACE_ZSH_HASH_DIRS="$(hash -d)"
  declare -x CARAPACE_SHELL=zsh
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
    print -l ${(k)aliases} > "$_carapace_state/aliases"
    print -roC1 -- ${(k)builtins} > "$_carapace_state/builtins"
    print -l ${(ok)functions} > "$_carapace_state/functions"
    print -l ${${(k)jobtexts}/#/%%%%} > "$_carapace_state/jobs"
    print -l ${(ok)parameters} > "$_carapace_state/variables"
  else
    unset CARAPACE_SHELL_STATE
  fi

  # shellcheck disable=SC2086,SC2154,SC2155
  lines="$(echo "${compline}''" | xargs carapace "${command}" zsh 2>/dev/null)"
  local _carapace_rc=$?
  if [ $_carapace_rc -eq 1 ]; then
    lines="$(echo "${compline}'" | xargs carapace "${command}" zsh 2>/dev/null)"
    _carapace_rc=$?
    if [ $_carapace_rc -eq 1 ]; then
      lines="$(echo "${compline}\"" | xargs carapace "${command}" zsh 2>/dev/null)"
      _carapace_rc=$?
    fi
  fi
  if [ $_carapace_rc -ne 0 ] && [ -n "${CARAPACE_SHELL_STATE:-}" ]; then
    local _carapace_old_state="${CARAPACE_SHELL_STATE}"
    unset CARAPACE_SHELL_STATE
    lines="$(echo "${compline}''" | xargs carapace "${command}" zsh 2>/dev/null)"
    _carapace_rc=$?
    if [ $_carapace_rc -eq 1 ]; then
      lines="$(echo "${compline}'" | xargs carapace "${command}" zsh 2>/dev/null)"
      _carapace_rc=$?
      if [ $_carapace_rc -eq 1 ]; then
        lines="$(echo "${compline}\"" | xargs carapace "${command}" zsh 2>/dev/null)"
      fi
    fi
    declare -x CARAPACE_SHELL_STATE="${_carapace_old_state}"
  fi

  local zstyle message noprefix data
  IFS=$'\001' read -r -d '' zstyle message noprefix data <<<"${lines}"
  # shellcheck disable=SC2154
  zstyle ":completion:${curcontext}:*" list-colors "${zstyle}"
  zstyle ":completion:${curcontext}:*" group-name ''
  [ -z "$message" ] || _message -r "${message}"
  [[ "${noprefix}" = "true" ]] && compstate[insert]=menu

  local block tag displays values displaysArr valuesArr
  while IFS=$'\002' read -r -d $'\002' block; do
    IFS=$'\003' read -r -d '' tag displays values <<<"${block}"
    # shellcheck disable=SC2034
    IFS=$'\n' read -r -d $'\004' -A displaysArr <<<"${displays}"$'\004'
    IFS=$'\n' read -r -d $'\004' -A valuesArr <<<"${values}"$'\004'

    [[ ${#valuesArr[@]} -gt 1 ]] && _describe -t "${tag}" "${tag}" displaysArr valuesArr -Q -S ''
  done <<<"${data}"
  [ -n "${_carapace_state_cleanup:-}" ] && rm -rf "$_carapace_state_cleanup"
}

compdef _carapace_completer example

