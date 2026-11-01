package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_files_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Create, update, or remove the configured system files and directories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_files_applyCmd).Standalone()

	bootstrap_files_applyCmd.Flags().BoolP("dry-run", "n", false, "Show what would change without changing anything")
	bootstrap_files_applyCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_files_applyCmd.Flags().Bool("prompt-secrets", false, "Prompt securely for missing bootstrap secret inputs")
	bootstrap_files_applyCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrap_filesCmd.AddCommand(bootstrap_files_applyCmd)
}
