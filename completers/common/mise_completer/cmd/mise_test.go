package cmd

import (
	"testing"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/sandbox"
)

func TestActivate(t *testing.T) {
	sandbox.Package(t, "github.com/carapace-sh/carapace-bin/completers/common/mise_completer")(func(s *sandbox.Sandbox) {
		s.Run("activate", "").
			Expect(carapace.ActionValues("bash", "elvish", "fish", "nu", "xonsh", "zsh", "pwsh", "powershell"))
	})
}

func TestTasksSort(t *testing.T) {
	sandbox.Package(t, "github.com/carapace-sh/carapace-bin/completers/common/mise_completer")(func(s *sandbox.Sandbox) {
		s.Run("tasks", "ls", "--sort", "").
			Expect(carapace.ActionValues("name", "alias", "description", "source").
				Usage("Column to sort by (default: name)"))

		s.Run("tasks", "ls", "--sort-order", "").
			Expect(carapace.ActionValues("asc", "desc").
				Usage("Sort direction (default: asc)"))
	})
}

func TestToolAlias(t *testing.T) {
	sandbox.Package(t, "github.com/carapace-sh/carapace-bin/completers/common/mise_completer")(func(s *sandbox.Sandbox) {
		s.Run("tool-alias", "").
			Expect(carapace.ActionValuesDescribed(
				"add", "Set a backend alias or a version alias in the global config",
				"create", "Set a backend alias or a version alias in the global config",
				"del", "Remove a backend alias or a version alias from the global config",
				"delete", "Remove a backend alias or a version alias from the global config",
				"get", "Show the version a tool's version alias stands for",
				"list", "List tool version aliases",
				"ls", "List tool version aliases",
				"remove", "Remove a backend alias or a version alias from the global config",
				"rm", "Remove a backend alias or a version alias from the global config",
				"set", "Set a backend alias or a version alias in the global config",
				"unset", "Remove a backend alias or a version alias from the global config",
			).Tag("commands"))
	})
}
