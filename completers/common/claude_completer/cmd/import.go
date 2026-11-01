package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import config from another AI coding agent into Claude Code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().Bool("dry-run", false, "Show what would be imported without writing anything")
	importCmd.Flags().BoolP("help", "h", false, "Display help for command")
	importCmd.Flags().Bool("yes", false, "Skip the interactive picker")
	rootCmd.AddCommand(importCmd)

	carapace.Gen(importCmd).PositionalCompletion(
		carapace.ActionValues("codex", "cursor", "gemini"),
	)
}
