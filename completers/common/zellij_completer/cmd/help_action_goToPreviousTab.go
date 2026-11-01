package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_goToPreviousTabCmd = &cobra.Command{
	Use:   "go-to-previous-tab",
	Short: "Go to the previous tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_goToPreviousTabCmd).Standalone()

	help_actionCmd.AddCommand(help_action_goToPreviousTabCmd)
}
