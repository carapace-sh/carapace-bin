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
	snippet := `%v%v

let carapace_completer = {|spans|
  # if the current command is an alias, get it's expansion
  let expanded_alias = (scope aliases | where name == $spans.0 | get -o 0.expansion)

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
| default { $carapace_completer } completer)

$env.config = $current
    `

	return fmt.Sprintf(snippet, pathSnippet("nushell"), envSnippet("nushell"), windowsSnippet, windowsSnippet)
}
