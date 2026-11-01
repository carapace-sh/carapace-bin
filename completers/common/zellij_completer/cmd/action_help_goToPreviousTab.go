package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_goToPreviousTabCmd = &cobra.Command{
	Use:   "go-to-previous-tab",
	Short: "Go to the previous tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_goToPreviousTabCmd).Standalone()

	action_helpCmd.AddCommand(action_help_goToPreviousTabCmd)
}
