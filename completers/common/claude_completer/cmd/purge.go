package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Delete all Claude Code state for a project (transcripts, tasks, file history, config entry)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(purgeCmd).Standalone()

	purgeCmd.Flags().Bool("all", false, "Purge state for every project (mutually exclusive with [path])")
	purgeCmd.Flags().Bool("dry-run", false, "List what would be deleted without deleting anything")
	purgeCmd.Flags().BoolP("help", "h", false, "Display help for command")
	purgeCmd.Flags().BoolP("interactive", "i", false, "Prompt for each item before deleting")
	purgeCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt")
	rootCmd.AddCommand(purgeCmd)

	carapace.Gen(purgeCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
