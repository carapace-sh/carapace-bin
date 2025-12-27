package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var revertCmd = &cobra.Command{
	Use:     "revert",
	Short:   "Revert one or more unit files to vendor version",
	GroupID: "unit file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revertCmd).Standalone()

	rootCmd.AddCommand(revertCmd)

	carapace.Gen(revertCmd).PositionalAnyCompletion(
		action.ActionUnits(revertCmd).FilterArgs(),
	)
}
