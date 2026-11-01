package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_origin_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Connect an origin repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_origin_setCmd).Standalone()

	dotfiles_origin_setCmd.Flags().String("branch", "", "The branch to sync (default: the repository's default branch)")
	dotfiles_origin_setCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_origin_setCmd.Flags().String("sync", "", "How to sync with the repository: sync, fetch-only, or manual")
	dotfiles_origin_setCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	dotfiles_originCmd.AddCommand(dotfiles_origin_setCmd)
}
