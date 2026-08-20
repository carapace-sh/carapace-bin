package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_detachCmd = &cobra.Command{
	Use:   "detach",
	Short: "Detach from the current session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_detachCmd).Standalone()

	help_actionCmd.AddCommand(help_action_detachCmd)
}
