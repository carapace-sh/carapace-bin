package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_goToTabCmd = &cobra.Command{
	Use:   "go-to-tab",
	Short: "Go to tab with index [index]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_goToTabCmd).Standalone()

	action_helpCmd.AddCommand(action_help_goToTabCmd)
}
