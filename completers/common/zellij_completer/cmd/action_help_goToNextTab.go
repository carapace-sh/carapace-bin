package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_goToNextTabCmd = &cobra.Command{
	Use:   "go-to-next-tab",
	Short: "Go to the next tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_goToNextTabCmd).Standalone()

	action_helpCmd.AddCommand(action_help_goToNextTabCmd)
}
