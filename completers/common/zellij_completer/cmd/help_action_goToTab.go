package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_goToTabCmd = &cobra.Command{
	Use:   "go-to-tab",
	Short: "Go to tab with index [index]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_goToTabCmd).Standalone()

	help_actionCmd.AddCommand(help_action_goToTabCmd)
}
