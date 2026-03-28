package snippet

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
