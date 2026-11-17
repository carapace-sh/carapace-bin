use builtin
# set paths = ['/tmp/carapace-snippet-fixtures/carapace/bin' $@paths]

put example | each {|c|
    set edit:completion:arg-completer[$c] = {|@arg|
        with [E:CARAPACE_SHELL = 'elvish'] [E:CARAPACE_SHELL_BUILTINS = (keys $builtin: | to-lines | slurp)] {
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
        }
    }
}

