package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_help_inlineCmd = &cobra.Command{
	Use:   "inline",
	Short: "Inline completion mode with small TUI overlay",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_help_inlineCmd).Standalone()

	ai_helpCmd.AddCommand(ai_help_inlineCmd)
}
