package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_detachCmd = &cobra.Command{
	Use:   "detach",
	Short: "Detach from the current session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_detachCmd).Standalone()

	action_helpCmd.AddCommand(action_help_detachCmd)
}
