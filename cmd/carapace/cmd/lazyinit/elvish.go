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

	var quotedCompleters []string
	for _, c := range completers {
		quotedCompleters = append(quotedCompleters, fmt.Sprintf("%q", c))
	}

	return fmt.Sprintf(snippet, pathSnippet("elvish"), strings.Join(quotedCompleters, " "), windowsSnippet)
}
