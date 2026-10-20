// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_dotfiles_includeCmd = &cobra.Command{
	Use:   "include",
	Short: "Stop excluding paths matching a glob",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_dotfiles_includeCmd).Standalone()

	bootstrap_dotfiles_includeCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_dotfilesCmd.AddCommand(bootstrap_dotfiles_includeCmd)
}
