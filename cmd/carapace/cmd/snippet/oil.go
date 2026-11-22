package snippet

import (
	"fmt"
	"slices"

	"github.com/carapace-sh/carapace-bin/pkg/env"
	shlex "github.com/carapace-sh/carapace-shlex/v2"
)

func Oil(completers []string) string {
	builtins := ""
	if slices.Contains(env.Builtins(), "oil") {
		builtins = `  local _carapace_state _carapace_state_cleanup
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
    alias | sed 's/alias //' | sed 's/=.*//' > "$_carapace_state/aliases"
    compgen -b > "$_carapace_state/builtins"
    compgen -A function > "$_carapace_state/functions"
    jobs 2>/dev/null | while read -r line; do [[ $line =~ \[([0-9]+)\] ]] && echo "%%${BASH_REMATCH[1]}"; done > "$_carapace_state/jobs"
    compgen -v > "$_carapace_state/variables"
  else
    unset CARAPACE_SHELL_STATE
  fi
`
	}

	return fmt.Sprintf(`%v%v

_carapace_completer() {
  declare -x CARAPACE_SHELL=oil
%v
  local command="${COMP_WORDS[0]}"
  local compline="${COMP_LINE:0:${COMP_POINT}}"
  local IFS=$'\n'
  mapfile -t COMPREPLY < <(echo "$compline" | sed -e "s/ \$/ ''/" -e 's/"/\"/g' | xargs carapace "${command}" oil 2>/dev/null)
  local _carapace_rc=$?
  if [ $_carapace_rc -ne 0 ] && [ -n "${CARAPACE_SHELL_STATE:-}" ]; then
    local _carapace_old_state="${CARAPACE_SHELL_STATE}"
    unset CARAPACE_SHELL_STATE
    mapfile -t COMPREPLY < <(echo "$compline" | sed -e "s/ \$/ ''/" -e 's/"/\"/g' | xargs carapace "${command}" oil 2>/dev/null)
    declare -x CARAPACE_SHELL_STATE="${_carapace_old_state}"
  fi
  [[ "${COMPREPLY[@]}" == "" ]] && COMPREPLY=() # fix for mapfile creating a non-empty array from empty command output
  [[ ${COMPREPLY[0]} == *[/=@:.,$'\001'] ]] && compopt -o nospace
  # TODO use mapfile
  # shellcheck disable=SC2206
  [[ ${#COMPREPLY[@]} -eq 1 ]] && COMPREPLY=(${COMPREPLY[@]%%%%$'\001'})
  [ -n "${_carapace_state_cleanup:-}" ] && rm -rf "$_carapace_state_cleanup"
}

complete -F _carapace_completer %v
`, pathSnippet("oil"), envSnippet("oil"), builtins, shlex.Join(completers, shlex.Oil))
}
