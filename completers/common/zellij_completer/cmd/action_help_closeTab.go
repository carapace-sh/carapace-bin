package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_closeTabCmd = &cobra.Command{
	Use:   "close-tab",
	Short: "Close the current tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_closeTabCmd).Standalone()

	action_helpCmd.AddCommand(action_help_closeTabCmd)
}
