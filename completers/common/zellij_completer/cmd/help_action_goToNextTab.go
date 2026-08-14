package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_goToNextTabCmd = &cobra.Command{
	Use:   "go-to-next-tab",
	Short: "Go to the next tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_goToNextTabCmd).Standalone()

	help_actionCmd.AddCommand(help_action_goToNextTabCmd)
}
