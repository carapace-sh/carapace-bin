package lazyinit

import (
	"fmt"
	"os"
	"runtime"
	"slices"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/xdg"
)

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
			snippet = fmt.Sprintf(`$env.Path = ($env.Path | split row (char esep) | where { $in != "%v" } | prepend "%v")`, fixedBinDir, fixedBinDir)
		} else {
			snippet = fmt.Sprintf(`$env.PATH = ($env.PATH | split row (char esep) | where { $in != "%v" } | prepend "%v")`, fixedBinDir, fixedBinDir)
		}

	case "powershell":
		snippet = fmt.Sprintf(`[Environment]::SetEnvironmentVariable("PATH", "%v" + [IO.Path]::PathSeparator + [Environment]::GetEnvironmentVariable("PATH"))`, binDir)

	case "xonsh":
		snippet = fmt.Sprintf(`__xonsh__.env['PATH'].insert(0, '%v')`, binDir)

	default:
		snippet = fmt.Sprintf("# error: unknown shell: %#v", shell)
	}

	if slices.Contains(strings.Split(os.Getenv("PATH"), string(os.PathListSeparator)), binDir) {
		carapace.LOG.Printf("PATH already contains %#v\n", binDir)
		if shell != "nushell" {
			snippet = "# " + snippet
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
