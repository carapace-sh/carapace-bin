package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_goToTabNameCmd = &cobra.Command{
	Use:   "go-to-tab-name",
	Short: "Go to tab with name [name]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_goToTabNameCmd).Standalone()

	action_helpCmd.AddCommand(action_help_goToTabNameCmd)
}
