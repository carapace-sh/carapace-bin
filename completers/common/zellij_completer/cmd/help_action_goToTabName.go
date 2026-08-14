package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_goToTabNameCmd = &cobra.Command{
	Use:   "go-to-tab-name",
	Short: "Go to tab with name [name]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_goToTabNameCmd).Standalone()

	help_actionCmd.AddCommand(help_action_goToTabNameCmd)
}
