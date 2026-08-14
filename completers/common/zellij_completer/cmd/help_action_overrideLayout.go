package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_overrideLayoutCmd = &cobra.Command{
	Use:   "override-layout",
	Short: "Override the layout of the active tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_overrideLayoutCmd).Standalone()

	help_actionCmd.AddCommand(help_action_overrideLayoutCmd)
}
