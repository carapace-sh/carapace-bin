// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_captureCmd = &cobra.Command{
	Use:   "capture",
	Short: "Run a command and save tracked files before and after it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_captureCmd).Standalone()

	dotfiles_captureCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfiles_captureCmd.Flags().String("label", "", "Describe this operation in history")
	dotfilesCmd.AddCommand(dotfiles_captureCmd)
}
