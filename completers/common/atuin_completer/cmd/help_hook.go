package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_hookCmd = &cobra.Command{
	Use:   "hook",
	Short: "Manage AI-agent shell hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_hookCmd).Standalone()

	helpCmd.AddCommand(help_hookCmd)
}
