package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/tlp_completer/cmd/action"
	"github.com/spf13/cobra"
)

var powerSaverCmd = &cobra.Command{
	Use:   "power-saver",
	Short: "Apply the power-saver profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(powerSaverCmd).Standalone()

	rootCmd.AddCommand(powerSaverCmd)

	carapace.Gen(powerSaverCmd).DashAnyCompletion(action.ActionConfigParams())
}
