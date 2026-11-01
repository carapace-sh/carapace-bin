package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var markUserCmd = &cobra.Command{
	Use:   "user [options] <package-spec>...",
	Short: "mark package as user-installed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(markUserCmd).Standalone()

	markCmd.AddCommand(markUserCmd)

	carapace.Gen(markUserCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(markUserCmd),
	)
}
