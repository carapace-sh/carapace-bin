package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var tryRestartCmd = &cobra.Command{
	Use:     "try-restart",
	Short:   "Restart one or more units if active",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tryRestartCmd).Standalone()

	rootCmd.AddCommand(tryRestartCmd)

	carapace.Gen(tryRestartCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits(tryRestartCmd).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
