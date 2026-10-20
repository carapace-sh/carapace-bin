// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show what `mise dot apply` would change",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_diffCmd).Standalone()

	bootstrap_dotfiles_diffCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_diffCmd.Flags().Bool("prompt-secrets", false, "Prompt for `[bootstrap.secrets]` values that templates need and the environment does not set")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_diffCmd)
}
