package cmd

//go:generate go run ../../carapace-generate/gen.go

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/xdg"
)

func bash_lazy(completers []string) string {
	snippet := `%v%v

_carapace_lazy() {
  source <(carapace $1 bash)
   $"_$1_completion"
}
complete -F _carapace_lazy %v
`
	return fmt.Sprintf(snippet, pathSnippet("bash"), envSnippet("bash"), strings.Join(completers, " "))
}

func bash_ble_lazy(completers []string) string {
	snippet := `%v

_carapace_lazy() {
  source <(carapace $1 bash-ble)
   $"_$1_completion_ble"
}
complete -F _carapace_lazy %v
`
	return fmt.Sprintf(snippet, pathSnippet("bash-ble"), strings.Join(completers, " "))
}

func elvish_lazy(completers []string) string {
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

func pathSnippet(shell string) (snippet string) {
	configDir, err := xdg.UserConfigDir()
	if err != nil {
		panic(err.Error())
	}
	binDir := configDir + "/carapace/bin"

	switch shell {
	case "bash", "bash-ble", "oil", "zsh":
		snippet = fmt.Sprintf(`export PATH="%v%v$PATH"`, binDir, string(os.PathListSeparator))

	case "elvish":
		snippet = fmt.Sprintf(`set paths = ['%v' $@paths]`, binDir)

	case "fish":
		snippet = fmt.Sprintf(`fish_add_path '%v'`, binDir)

	case "nushell":
		fixedBinDir := strings.ReplaceAll(binDir, `\`, `\\`)
		if runtime.GOOS == "windows" {
			snippet = fmt.Sprintf(`$env.Path = ($env.Path | split row (char esep) | append "%v")`, fixedBinDir)
		} else {
			snippet = fmt.Sprintf(`$env.PATH = ($env.PATH | split row (char esep) | append "%v")`, fixedBinDir)
		}

	case "powershell":
		snippet = fmt.Sprintf(`[Environment]::SetEnvironmentVariable("PATH", "%v" + [IO.Path]::PathSeparator + [Environment]::GetEnvironmentVariable("PATH"))`, binDir)

	case "xonsh":
		snippet = fmt.Sprintf(`__xonsh__.env['PATH'].insert(0, '%v')`, binDir)

	default:
		snippet = fmt.Sprintf("# error: unknown shell: %#v", shell)
	}

	for _, path := range strings.Split(os.Getenv("PATH"), string(os.PathListSeparator)) {
		if path == binDir {
			carapace.LOG.Printf("PATH already contains %#v\n", binDir)
			if shell != "nushell" {
				snippet = "# " + snippet
			}
			break
		}
	}
	return
}

func envSnippet(shell string) string {
	if os.Getenv("CARAPACE_ENV") == "0" {
		return ""
	}

	switch shell {
	case "bash", "oil":
		return `

get-env () { echo "${!1}"; }
set-env () { export "$1=$2"; }
unset-env () { unset "$1"; }`

	case "fish":
		return `

function get-env -d "get environment variable"; echo $$argv[1]; end
function set-env -d "set environment variable"; set -g -x $argv[1] $argv[2]; end
function unset-env -d "unset environment variable"; set -e $argv[1]; end`

	case "nushell":
		return `

def --env get-env [name] { $env | get $name }
def --env set-env [name, value] { load-env { $name: $value } }
def --env unset-env [name] { hide-env $name }`

	case "powershell":
		return `

Function get-env([string]$name) { Get-Item "env:$name" }
Function set-env([string]$name, [string]$value) { Set-Item "env:$name" "$value" }
Function unset-env([string]$name) { Remove-Item "env:$name" }`

	case "xonsh":
		return `

def _carapace_getenv(args):
	print(__xonsh__.env[args[0]])
aliases['get-env']=_carapace_getenv

def _carapace_setenv(args):
	__xonsh__.env[args[0]]=args[1]
	os.environ[args[0]]=args[1]
aliases['set-env']=_carapace_setenv

def _carapace_unsetenv(args):
	del __xonsh__.env[args[0]]
	del os.environ[args[0]]
aliases['unset-env']=_carapace_unsetenv`

	case "zsh":
		return `

get-env () { echo "${(P)1}"; }
set-env () { export "$1=$2"; }
unset-env () { unset "$1"; }`

	default:
		return ""
	}
}

func fish_lazy(completers []string) string {
	snippet := `%v%v

function _carapace_lazy
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
	return fmt.Sprintf(snippet, pathSnippet("fish"), envSnippet("fish"), strings.Join(complete, "\n"))
}

func nushell_lazy(completers []string) string {
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

func oil_lazy(completers []string) string {
	snippet := `%v%v

_carapace_lazy() {
  source <(carapace $1 oil)
   $"_$1_completion"
}
complete -F _carapace_lazy %v
`
	return fmt.Sprintf(snippet, pathSnippet("oil"), envSnippet("oil"), strings.Join(completers, " "))
}

func powershell_lazy(completers []string) string {
	snippet := `%v%v

$_carapace_lazy = {
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
	return fmt.Sprintf(snippet, pathSnippet("powershell"), envSnippet("powershell"), strings.Join(complete, "\n"))
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
import os

%v%v

@contextual_completer
def _carapace_lazy(context):
    """carapace lazy"""
    if (context.command and
        context.command.arg_index > 0 and
        context.command.args[0].value in [%v]):
        exec(compile(subprocess.run(['carapace', context.command.args[0].value, 'xonsh'], stdout=subprocess.PIPE).stdout.decode('utf-8'), "", "exec"))
        return XSH.completers[context.command.args[0].value](context)
`
	complete := make([]string, len(completers))
	for index, completer := range completers {
		complete[index] = fmt.Sprintf(`'%v'`, completer)
	}
	snippet += `_add_one_completer('carapace_lazy', _carapace_lazy, 'start')`
	return fmt.Sprintf(snippet, pathSnippet("xonsh"), envSnippet("xonsh"), strings.Join(complete, ", "))
}

func zsh_lazy(completers []string) string {
	snippet := `%v%v

function _carapace_lazy {
    source <(carapace $words[1] zsh)
}
compdef _carapace_lazy %v
`
	return fmt.Sprintf(snippet, pathSnippet("zsh"), envSnippet("zsh"), strings.Join(completers, " "))
}
