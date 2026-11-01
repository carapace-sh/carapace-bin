package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/tlp_completer/cmd/action"
	"github.com/spf13/cobra"
)

var batCmd = &cobra.Command{
	Use:   "bat",
	Short: "Apply profile for battery power and enter manual mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batCmd).Standalone()

	rootCmd.AddCommand(batCmd)

	carapace.Gen(batCmd).DashAnyCompletion(action.ActionConfigParams())
}
