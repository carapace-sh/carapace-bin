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

# nushell passes a place record since 0.116, the element text since 0.115 and only the token before that
def _carapace-env-span [command: string, place: any] {
  let described = ($place | describe)
  let parts = (if $described =~ "record" {
      ($place.command? | default [])
    } else if $described =~ "list" {
      $place
    } else {
      ($place | split row " ")
    })

  if ($parts | is-empty) {
    [$command, ""]
  } else if (($parts | length) > 1) and (($parts | first) == $command) {
    $parts
  } else {
    [$command, ...$parts]
  }
}

def _carapace-env-complete [command: string, place: any] {
  let spans = (_carapace-env-span $command $place)

  with-env { CARAPACE_SHELL: "nushell" } {
    carapace $spans.0 nushell ...$spans | from json
  }
}

def _carapace-get-env [place?] { _carapace-env-complete get-env $place }
def _carapace-set-env [place?] { _carapace-env-complete set-env $place }
def _carapace-unset-env [place?] { _carapace-env-complete unset-env $place }

def --env get-env [name: any@_carapace-get-env] { $env | get $name }
def --env set-env [name: any@_carapace-set-env, value: any@_carapace-set-env] { load-env { $name: $value } }
def --env unset-env [name: any@_carapace-unset-env] { hide-env $name }`

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
