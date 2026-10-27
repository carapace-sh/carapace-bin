package snippet

import (
	"fmt"
	"os"
	"runtime"
	"slices"
	"strings"

	"github.com/carapace-sh/carapace"
	shlex "github.com/carapace-sh/carapace-shlex/v2"
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
		snippet = fmt.Sprintf(`export PATH="%v%v$PATH"`, shlex.Escape(binDir, shlex.Bash), string(os.PathListSeparator))

	case "elvish":
		snippet = fmt.Sprintf(`set paths = ['%v' $@paths]`, shlex.Escape(binDir, shlex.Elvish))

	case "fish":
		snippet = fmt.Sprintf(`fish_add_path "%v"`, shlex.Escape(binDir, shlex.Fish))

	case "nushell":
		fixedBinDir := strings.NewReplacer(`\`, `\\`, `"`, `\"`).Replace(binDir)
		if runtime.GOOS == "windows" {
			snippet = fmt.Sprintf(`$env.Path = ($env.Path | split row (char esep) | where { $in != "%v" } | prepend "%v")`, fixedBinDir, fixedBinDir)
		} else {
			snippet = fmt.Sprintf(`$env.PATH = ($env.PATH | split row (char esep) | where { $in != "%v" } | prepend "%v")`, fixedBinDir, fixedBinDir)
		}

	case "powershell":
		snippet = fmt.Sprintf(`[Environment]::SetEnvironmentVariable("PATH", '%v' + [IO.Path]::PathSeparator + [Environment]::GetEnvironmentVariable("PATH"))`, shlex.Escape(binDir, shlex.Powershell))

	case "xonsh":
		snippet = fmt.Sprintf(`__xonsh__.env['PATH'].insert(0, '%v')`, shlex.Escape(binDir, shlex.Xonsh))

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
