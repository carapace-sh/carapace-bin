package lazyinit

import "fmt"

func Nushell(completers []string) string {
	snippet := `%v%v

let carapace_completer = {|spans| 
  carapace $spans.0 nushell $spans | from json
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
