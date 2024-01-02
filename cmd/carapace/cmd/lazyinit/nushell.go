package lazyinit

import "fmt"

func Nushell(completers []string) string {
	snippet := `%v%v

let carapace_completer = {|spans|
  # if the current command is an alias, get it's expansion
  let expanded_alias = (scope aliases | where name == $spans.0 | get -i 0 | get -i expansion)

  # overwrite
  let spans = (if $expanded_alias != null  {
    # put the first word of the expanded alias first in the span
    $spans | skip 1 | prepend ($expanded_alias | split row " ")
  } else {
    $spans
  })

  carapace $spans.0 nushell $spans
  | from json
  | if ($in | default [] | where value =~ '^-.*ERR$' | is-empty) { $in } else { null }
}

mut current = (($env | default {} config).config | default {} completions)
$current.completions = ($current.completions | default {} external)
$current.completions.external = ($current.completions.external
    | default true enable
    | default $carapace_completer completer)

$env.config = $current
    `

	return fmt.Sprintf(snippet, pathSnippet("nushell"), envSnippet("nushell"))
}
