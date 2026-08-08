package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var markDependencyCmd = &cobra.Command{
	Use:   "dependency [options] <package-spec>...",
	Short: "mark package as a dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(markDependencyCmd).Standalone()

	markCmd.AddCommand(markDependencyCmd)

	carapace.Gen(markDependencyCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(markDependencyCmd),
	)
}
