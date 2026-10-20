// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_origin_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Connect an origin repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_origin_setCmd).Standalone()

	bootstrap_dotfiles_origin_setCmd.Flags().String("branch", "", "The branch to sync (default: the repository's default branch)")
	bootstrap_dotfiles_origin_setCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_origin_setCmd.Flags().String("sync", "", "How to sync with the repository: sync, fetch-only, or manual")
	bootstrap_dotfiles_origin_setCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrap_dotfiles_originCmd.AddCommand(bootstrap_dotfiles_origin_setCmd)
}
