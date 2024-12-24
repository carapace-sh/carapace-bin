package lazyinit

import (
	"fmt"
	"runtime"
	"strings"
)

func Elvish(completers []string) string {
	windowsSnippet := ""
	if runtime.GOOS == "windows" {
		windowsSnippet = "\n    set edit:completion:arg-completer[$c.exe] = $edit:completion:arg-completer[$c]\n"
	}

	snippet := `%v

put %v | each {|c|
    set edit:completion:arg-completer[$c] = {|@arg|
        set edit:completion:arg-completer[$c] = {|@arg| }
        eval (carapace $c elvish | slurp)
        $edit:completion:arg-completer[$c] $@arg
    }%v
}
`
	return fmt.Sprintf(snippet, pathSnippet("elvish"), strings.Join(completers, " "), windowsSnippet)
}
