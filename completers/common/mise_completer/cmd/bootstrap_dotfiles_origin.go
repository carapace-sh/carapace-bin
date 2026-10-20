// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_originCmd = &cobra.Command{
	Use:   "origin",
	Short: "Show, connect, or disconnect the origin repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_originCmd).Standalone()

	bootstrap_dotfiles_originCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfiles_originCmd.Flags().Bool("remove", false, "Disconnect the origin by removing `[history.origin]`; local checkpoints and fetched history stay")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_originCmd)
}
