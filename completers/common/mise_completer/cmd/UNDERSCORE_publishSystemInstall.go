// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var __publishSystemInstallCmd = &cobra.Command{
	Use:    "__publish-system-install",
	Short:  "[internal] publish a staged system installation as root",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(__publishSystemInstallCmd).Standalone()

	__publishSystemInstallCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(__publishSystemInstallCmd)
}
