package snippet

import (
	"os"
)

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
