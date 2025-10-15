package lazyinit

import (
	"fmt"
	"runtime"
)

func Nushell(completers []string) string {
	windowsSnippet := ""
	if runtime.GOOS == "windows" {
		windowsSnippet = " | str replace --regex  '\\.exe$' ''"
	}
	snippet := `export-env {
  %v

  let carapace_completer = {|spans|
    # if the current command is an alias, get it's expansion
    let expanded_alias = (scope aliases | where name == $spans.0 | $in.0?.expansion?)

    # overwrite
    let spans = (if $expanded_alias != null  {
      # put the first word of the expanded alias first in the span
      $spans | skip 1 | prepend ($expanded_alias | split row " " | take 1%v)
    } else {
      $spans | skip 1 | prepend ($spans.0%v)
    })

    carapace $spans.0 nushell ...$spans
    | from json
  }

  mut current = (($env | default {} config).config | default {} completions)
  $current.completions = ($current.completions | default {} external)
  $current.completions.external = ($current.completions.external
  | default true enable
  # backwards compatible workaround for default, see nushell #15654
  | upsert completer { if $in == null { $carapace_completer } else { $in } })

  $env.config = $current
}%v`

	return fmt.Sprintf(snippet, pathSnippet("nushell"), windowsSnippet, windowsSnippet, envSnippet("nushell"))
}
