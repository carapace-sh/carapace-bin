package lazyinit

import (
	"fmt"
	"strings"
)

func Elvish(completers []string) string {
	snippet := `%v

put %v | each {|c|
    set edit:completion:arg-completer[$c] = {|@arg|
        set edit:completion:arg-completer[$c] = {|@arg| }
        eval (carapace $c elvish | slurp)
        $edit:completion:arg-completer[$c] $@arg
    }
}
`
	return fmt.Sprintf(snippet, pathSnippet("elvish"), strings.Join(completers, " "))
}
