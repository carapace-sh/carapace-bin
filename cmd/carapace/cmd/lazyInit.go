package cmd

//go:generate go run ../../generate/gen.go

import (
	"fmt"
	"os"
	"strings"

	"github.com/mitchellh/go-ps"
)

func bash_lazy(completers []string) string {
	snippet := `_carapace_lazy() {
  source <(carapace $1 bash)
   $"_$1_completions"
}
complete -F _carapace_lazy %v
`
	return fmt.Sprintf(snippet, strings.Join(completers, " "))
}

func elvish_lazy(completers []string) string {
	snippet := `put %v | each [c]{
    edit:completion:arg-completer[$c] = [@arg]{
        edit:completion:arg-completer[$c] = [@arg]{}
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
		complete[index] = fmt.Sprintf(` complete --erase -c '%v'
complete -c '%v' -f -a '(_carapace_lazy %v)'`, completer, completer, completer)
	}
	return fmt.Sprintf(snippet, strings.Join(complete, "\n"))
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
    & (Get-Variable -name "_${completer}_completer").Value $wordToComplete $commandAst $cursorPosition
}
%v
`
	complete := make([]string, len(completers))
	for index, completer := range completers {
		complete[index] = fmt.Sprintf(`Register-ArgumentCompleter -Native -CommandName '%v' -ScriptBlock $_carapace_lazy`, completer)
	}
	return fmt.Sprintf(snippet, strings.Join(complete, "\n"))
}

func xonsh_lazy(completers []string) string {
	snippet := `from shlex import split
import re
import pathlib
import subprocess
import xonsh
import builtins
from xonsh.completers._aliases import _add_one_completer
from xonsh.completers.path import complete_dir, complete_path
from xonsh.completers.tools import RichCompletion

def _carapace_lazy(prefix, line, begidx, endidx, ctx):
    """carapace lazy"""
    command = split(line)[0]
    if command not in [%v]:
        return # not the expected command to complete
    builtins.__xonsh__.completers = builtins.__xonsh__.completers.copy()
    exec(compile(subprocess.run(['carapace', command, 'xonsh'], stdout=subprocess.PIPE).stdout.decode('utf-8'), "", "exec"))
    return builtins.__xonsh__.completers[command](prefix, line, begidx, endidx, ctx)
_add_one_completer('carapace_lazy', _carapace_lazy, 'start')
`
	complete := make([]string, len(completers))
	for index, completer := range completers {
		complete[index] = fmt.Sprintf(`'%v'`, completer)
	}
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

func determineShell() string {
	process, _ := ps.FindProcess(os.Getpid())
	for {
		if parentProcess, err := ps.FindProcess(process.PPid()); err != nil {
			return ""
		} else {
			switch parentProcess.Executable() {
			case "bash":
				return "bash"
			case "elvish":
				return "elvish"
			case "fish":
				return "fish"
			case "osh":
				return "oil"
			case "pwsh":
				return "powershell"
			case "xonsh":
				return "xonsh"
			case "zsh":
				return "zsh"
			default:
				return ""
			}
		}
	}
}
