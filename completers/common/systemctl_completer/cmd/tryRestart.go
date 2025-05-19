package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemctl_completer/cmd/action"
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
		action.ActionUnits(tryRestartCmd).FilterArgs(),
	)
}
