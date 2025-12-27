package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:     "kill",
	Short:   "Send signal to processes of a unit",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killCmd).Standalone()

	rootCmd.AddCommand(killCmd)

	carapace.Gen(killCmd).PositionalAnyCompletion(
		action.ActionUnits(killCmd).FilterArgs(),
	)
}
