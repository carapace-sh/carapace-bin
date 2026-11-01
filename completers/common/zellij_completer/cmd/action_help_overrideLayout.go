package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_overrideLayoutCmd = &cobra.Command{
	Use:   "override-layout",
	Short: "Override the layout of the active tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_overrideLayoutCmd).Standalone()

	action_helpCmd.AddCommand(action_help_overrideLayoutCmd)
}
