package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
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
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits(revertCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
