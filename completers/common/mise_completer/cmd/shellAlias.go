// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shellAliasCmd = &cobra.Command{
	Use:   "shell-alias",
	Short: "Manage shell aliases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shellAliasCmd).Standalone()

	shellAliasCmd.Flags().BoolP("help", "h", false, "Print help")
	shellAliasCmd.Flags().Bool("no-header", false, "Do not print the table header")
	rootCmd.AddCommand(shellAliasCmd)
}
