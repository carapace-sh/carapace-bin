package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ai_inlineCmd = &cobra.Command{
	Use:   "inline",
	Short: "Inline completion mode with small TUI overlay",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ai_inlineCmd).Standalone()

	help_aiCmd.AddCommand(help_ai_inlineCmd)
}
