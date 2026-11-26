package snippet

import (
	"fmt"
	"runtime"
	"slices"

	"github.com/carapace-sh/carapace-bin/pkg/env"
	shlex "github.com/carapace-sh/carapace-shlex/v2"
)

func Elvish(completers []string) string {
	windowsSnippet := ""
	if runtime.GOOS == "windows" {
		windowsSnippet = "\n    set edit:completion:arg-completer[$c.exe] = $edit:completion:arg-completer[$c]\n"
	}

	envVars := "[E:CARAPACE_SHELL = 'elvish']"
	pathImport := ""
	if slices.Contains(env.Builtins(), "elvish") {
		pathImport = "\nuse path"
		envVars = `[E:CARAPACE_SHELL = 'elvish'] [E:CARAPACE_SHELL_STATE = (try { var dir = (if (and (has-env XDG_RUNTIME_DIR) (path:is-dir $E:XDG_RUNTIME_DIR)) { path:join $E:XDG_RUNTIME_DIR carapace builtins (to-string (num (id -u))) (to-string (num $pid)) } else { mktemp -d (path:join (if (has-env TMPDIR) { get-env TMPDIR } else { put /tmp }) carapace-builtin.XXXXXX) }); mkdir -p -m 700 $dir; print (keys $builtin: | to-lines | slurp) > (path:join $dir builtins); put $dir } catch { put '' })]`
	}

	snippet := `use builtin
%v%v

put %v | each {|c|
    set edit:completion:arg-completer[$c] = {|@arg|
        with %v {
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
}
`

	return fmt.Sprintf(snippet, pathSnippet("elvish"), pathImport, shlex.Join(completers, shlex.Elvish), envVars, windowsSnippet)
}
