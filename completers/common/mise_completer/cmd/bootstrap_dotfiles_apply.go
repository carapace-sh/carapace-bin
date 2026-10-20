// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Link, copy, and render `[dotfiles]` entries into place",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_applyCmd).Standalone()

	bootstrap_dotfiles_applyCmd.Flags().BoolP("dry-run", "n", false, "Print the actions that would run without writing anything")
	bootstrap_dotfiles_applyCmd.Flags().BoolP("force", "f", false, "Overwrite existing files that conflict with whole-file dotfile entries")
	bootstrap_dotfiles_applyCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_applyCmd.Flags().Bool("prompt-secrets", false, "Prompt for `[bootstrap.secrets]` values that templates need and the environment does not set")
	bootstrap_dotfiles_applyCmd.Flags().Bool("prune", false, "Also remove files that dotfile groups deployed but no active entry deploys now")
	bootstrap_dotfiles_applyCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_applyCmd)
}
