// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_unapplyCmd = &cobra.Command{
	Use:   "unapply",
	Short: "Remove files that `mise dot apply` deployed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_unapplyCmd).Standalone()

	dotfiles_unapplyCmd.Flags().BoolP("dry-run", "n", false, "Print the actions that would run without writing anything")
	dotfiles_unapplyCmd.Flags().BoolP("force", "f", false, "Remove modified or otherwise ambiguous managed files and lines")
	dotfiles_unapplyCmd.Flags().StringSlice("group", nil, "Only unapply the files of this dotfile group, even one that is no longer selected or declared")
	dotfiles_unapplyCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_unapplyCmd.Flags().Bool("prompt-secrets", false, "Prompt for `[bootstrap.secrets]` values that templates need and the environment does not set")
	dotfiles_unapplyCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	dotfilesCmd.AddCommand(dotfiles_unapplyCmd)
}
