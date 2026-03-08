package mise

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/carapace-sh/carapace/pkg/style"
)

var completer = carapace.Gen(new(Cmd))

type Cmd struct{}

func (c *Cmd) Execute(args ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionCommands(
			"activate", "alias", "bin-paths", "config", "current", "deactivate", "doctor", "env", "exec", "generate", "hook-env", "install", "ls", "ls-remote", "outdated", "plugins", "prune", "quite", "registry", "run", "self-update", "settings", "shell", "tasks", "trust", "uninstall", "use", "version", "watch", "where", "which",
		).Invoke(c).ToAction().Usage("manage tools").Style(style.Blue)
	}).PositionalAny(carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch c.Args[0] {
		case "use", "install", "uninstall", "shell", "run", "tasks", "where", "which", "exec":
			return mise.ActionToolVersions()
		default:
			return carapace.ActionValues()
		}
	}))
}
