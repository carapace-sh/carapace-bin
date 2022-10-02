package cmd

//go:generate go run ../../generate/gen.go

import (
	"fmt"
	"strings"
)

func bash_lazy(completers []string) string {
	snippet := `_carapace_lazy() {
  source <(carapace $1 bash)
   $"_$1_completion"
}
complete -F _carapace_lazy %v
`
	return fmt.Sprintf(snippet, strings.Join(completers, " "))
}

func bash_ble_lazy(completers []string) string {
	snippet := `_carapace_lazy() {
  source <(carapace $1 bash-ble)
   $"_$1_completion_ble"
}
complete -F _carapace_lazy %v
`
	return fmt.Sprintf(snippet, strings.Join(completers, " "))
}

func elvish_lazy(completers []string) string {
	snippet := `put %v | each {|c|
    set edit:completion:arg-completer[$c] = {|@arg|
        set edit:completion:arg-completer[$c] = {|@arg| }
        eval (carapace $c elvish | slurp)
        $edit:completion:arg-completer[$c] $@arg
    }
}
`
	return fmt.Sprintf(snippet, strings.Join(completers, " "))
}

func fish_lazy(completers []string) string {
	snippet := `function _carapace_lazy
   complete -c $argv[1] -e
   carapace $argv[1] fish | source
   complete --do-complete=(commandline -cp)
end
%v
`
	complete := make([]string, len(completers))
	for index, completer := range completers {
		complete[index] = fmt.Sprintf(`complete -c '%v' -f -a '(_carapace_lazy %v)'`, completer, completer)
	}
	return fmt.Sprintf(snippet, strings.Join(complete, "\n"))
}

func nushell_lazy(completers []string) string {
	return `let carapace_completer = {|spans| 
  carapace $spans.0 nushell $spans | from json
}

let-env config = {
  external_completer: $carapace_completer
}`
}

func oil_lazy(completers []string) string {
	snippet := `_carapace_lazy() {
  source <(carapace $1 oil)
   $"_$1_completion"
}
complete -F _carapace_lazy %v
`
	return fmt.Sprintf(snippet, strings.Join(completers, " "))
}

func powershell_lazy(completers []string) string {
	snippet := `$_carapace_lazy = {
    param($wordToComplete, $commandAst, $cursorPosition)
    $completer = $commandAst.CommandElements[0].Value
    carapace $completer powershell | Out-String | Invoke-Expression
    & (Get-Item "Function:_${completer}_completer") $wordToComplete $commandAst $cursorPosition
}
%v
`
	complete := make([]string, len(completers))
	for index, completer := range completers {
		complete[index] = fmt.Sprintf(`Register-ArgumentCompleter -Native -CommandName '%v' -ScriptBlock $_carapace_lazy`, completer)
	}
	return fmt.Sprintf(snippet, strings.Join(complete, "\n"))
}

func tcsh_lazy(completers []string) string {
	// TODO hardcoded for now
	snippet := make([]string, len(completers))
	for index, c := range completers {
		snippet[index] = fmt.Sprintf("complete \"%v\" 'p@*@`echo \"$COMMAND_LINE'\"''\"'\" | xargs carapace %v tcsh `@@' ;", c, c)
	}
	return strings.Join(snippet, "\n")
}

func xonsh_lazy(completers []string) string {
	snippet := `from xonsh.completers._aliases import _add_one_completer
from xonsh.completers.tools import *

@contextual_completer
def _carapace_lazy(context):
    """carapace lazy"""
    if (context.command and
        context.command.arg_index > 0 and
        context.command.args[0].value in [%v]):
        XSH.completers = XSH.completers.copy()
        exec(compile(subprocess.run(['carapace', context.command.args[0].value, 'xonsh'], stdout=subprocess.PIPE).stdout.decode('utf-8'), "", "exec"))
        return XSH.completers[context.command.args[0].value](context)
`
	complete := make([]string, len(completers))
	for index, completer := range completers {
		complete[index] = fmt.Sprintf(`'%v'`, completer)
	}
	snippet += `_add_one_completer('carapace_lazy', _carapace_lazy, 'start')`
	return fmt.Sprintf(snippet, strings.Join(complete, ", "))
}

func zsh_lazy(completers []string) string {
	snippet := `function _carapace_lazy {
    source <(carapace $words[1] zsh)
}
compdef _carapace_lazy %v
`
	return fmt.Sprintf(snippet, strings.Join(completers, " "))
}
