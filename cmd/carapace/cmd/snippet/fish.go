package snippet

import (
	"fmt"
	"slices"
	"strings"

	"github.com/carapace-sh/carapace-bin/pkg/env"
	shlex "github.com/carapace-sh/carapace-shlex/v2"
)

func Fish(completers []string) string {
	builtins := ""
	if slices.Contains(env.Builtins(), "fish") {
		builtins = `  set --local _carapace_state ""
  set --local _carapace_state_cleanup ""
  if set -q XDG_RUNTIME_DIR; and test -d "$XDG_RUNTIME_DIR"; and test -O "$XDG_RUNTIME_DIR"; and not test -L "$XDG_RUNTIME_DIR"
    set _carapace_state "$XDG_RUNTIME_DIR/carapace/builtins/"(id -u)"/"(string replace -r '.*\\.' '' -- %self)
  else
    set _carapace_state (mktemp -d (string join '' (set -q TMPDIR; and echo $TMPDIR; or echo /tmp) /carapace-builtin.XXXXXX) 2>/dev/null; or true)
    set _carapace_state_cleanup "$_carapace_state"
  end
  if test -n "$_carapace_state"; and mkdir -m 700 -p "$_carapace_state" 2>/dev/null; and test -O "$_carapace_state"; and not test -L "$_carapace_state"
    set --local --export CARAPACE_SHELL_STATE "$_carapace_state"
    alias | string replace -r 'alias ' '' | string replace -r '=.*' '' > "$_carapace_state/aliases"
    builtin --names > "$_carapace_state/builtins"
    functions --all > "$_carapace_state/functions"
    jobs 2>/dev/null | string match -r '^\d+\s+\d+\s+(running|stopped|suspended|continued)' | string replace -r '^(\d+).*' '%%$1' > "$_carapace_state/jobs"
    set --names > "$_carapace_state/variables"
  else
    set -e CARAPACE_SHELL_STATE
  end
`
	}

	snippet := `%v%v

function _carapace_completer
  set --local --export CARAPACE_SHELL fish
%v  set --local data
  IFS='' set data (echo (commandline -cp)'' | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
  set --local _carapace_rc $status
  if test $_carapace_rc -eq 1
    IFS='' set data (echo (commandline -cp)"'" | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
    set _carapace_rc $status
    if test $_carapace_rc -eq 1
      IFS='' set data (echo (commandline -cp)'"' | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
      set _carapace_rc $status
    end
  end
  if test $_carapace_rc -ne 0; and set -q CARAPACE_SHELL_STATE
    set --local _carapace_old_state "$CARAPACE_SHELL_STATE"
    set -e CARAPACE_SHELL_STATE
    IFS='' set data (echo (commandline -cp)'' | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
    set _carapace_rc $status
    if test $_carapace_rc -eq 1
      IFS='' set data (echo (commandline -cp)"'" | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
      set _carapace_rc $status
      if test $_carapace_rc -eq 1
        IFS='' set data (echo (commandline -cp)'"' | sed "s/ \$/ ''/" | xargs carapace $argv[1] fish 2>/dev/null)
      end
    end
    set --export CARAPACE_SHELL_STATE "$_carapace_old_state"
  end
  if test -n "$_carapace_state_cleanup"
    rm -rf "$_carapace_state_cleanup"
  end
  echo $data
end

%v
`
	complete := make([]string, 0, len(completers)*2)
	for _, completer := range completers {
		name := shlex.Quote(completer, shlex.Fish)
		complete = append(complete,
			fmt.Sprintf(`complete -e %v`, name),
			fmt.Sprintf(`complete -c %v -f -a '(_carapace_completer %v)'`, name, name),
		)
	}
	return fmt.Sprintf(snippet, pathSnippet("fish"), envSnippet("fish"), builtins, strings.Join(complete, "\n"))
}
