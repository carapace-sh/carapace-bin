package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Link, copy, and render `[dotfiles]` entries into place",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_applyCmd).Standalone()

	dotfiles_applyCmd.Flags().BoolP("dry-run", "n", false, "Print the actions that would run without writing anything")
	dotfiles_applyCmd.Flags().BoolP("force", "f", false, "Overwrite existing files that conflict with whole-file dotfile entries")
	dotfiles_applyCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_applyCmd.Flags().Bool("prompt-secrets", false, "Prompt for `[bootstrap.secrets]` values that templates need and the environment does not set")
	dotfiles_applyCmd.Flags().Bool("prune", false, "Also remove files that dotfile groups deployed but no active entry deploys now")
	dotfiles_applyCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	dotfilesCmd.AddCommand(dotfiles_applyCmd)
}
