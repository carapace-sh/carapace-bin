// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_recoverCmd = &cobra.Command{
	Use:   "recover",
	Short: "Recover an interrupted dotfile operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_recoverCmd).Standalone()

	dotfiles_recoverCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_recoverCmd.Flags().Bool("keep-current", false, "Keep the files on disk instead of restoring the recovery copies")
	dotfiles_recoverCmd.Flags().BoolP("yes", "y", false, "Discard the operation's recovery copies without prompting")
	dotfilesCmd.AddCommand(dotfiles_recoverCmd)
}
