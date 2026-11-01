package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_newTabCmd = &cobra.Command{
	Use:   "new-tab",
	Short: "Create a new tab, optionally with a specified tab layout and name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_newTabCmd).Standalone()

	action_helpCmd.AddCommand(action_help_newTabCmd)
}
