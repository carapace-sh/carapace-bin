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
        carapace $c elvish (all $arg) | from-json | each {|completion|
    		put $completion[Messages] | all (one) | each {|m|
    			edit:notify (styled "error: " red)$m
    		}
    		if (not-eq $completion[Usage] "") {
    			edit:notify (styled "usage: " $completion[DescriptionStyle])$completion[Usage]
    		}
    		put $completion[Candidates] | all (one) | peach {|c|
    			if (eq $c[Description] "") {
    		    	edit:complex-candidate $c[Value] &display=(styled $c[Display] $c[Style]) &code-suffix=$c[CodeSuffix]
    			} else {
    		    	edit:complex-candidate $c[Value] &display=(styled $c[Display] $c[Style])(styled " " $completion[DescriptionStyle]" bg-default")(styled "("$c[Description]")" $completion[DescriptionStyle]) &code-suffix=$c[CodeSuffix]
    			}
    		}
        }
    }%v
}
`

	var quotedCompleters []string
	for _, c := range completers {
		quotedCompleters = append(quotedCompleters, fmt.Sprintf("%q", c))
	}

	return fmt.Sprintf(snippet, pathSnippet("elvish"), strings.Join(quotedCompleters, " "), windowsSnippet)
}
