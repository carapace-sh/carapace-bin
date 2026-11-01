package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show what `mise dot apply` would change",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_diffCmd).Standalone()

	dotfiles_diffCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_diffCmd.Flags().Bool("prompt-secrets", false, "Prompt for `[bootstrap.secrets]` values that templates need and the environment does not set")
	dotfilesCmd.AddCommand(dotfiles_diffCmd)
}
