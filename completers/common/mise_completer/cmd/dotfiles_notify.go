// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dotfiles_notifyCmd = &cobra.Command{
	Use:   "notify",
	Short: "Send a test desktop notification",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dotfiles_notifyCmd).Standalone()

	dotfiles_notifyCmd.Flags().BoolP("help", "h", false, "Print help")
	dotfilesCmd.AddCommand(dotfiles_notifyCmd)
}
