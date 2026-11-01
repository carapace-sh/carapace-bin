package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_newTabCmd = &cobra.Command{
	Use:   "new-tab",
	Short: "Create a new tab, optionally with a specified tab layout and name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_newTabCmd).Standalone()

	help_actionCmd.AddCommand(help_action_newTabCmd)
}
