// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Print the script to disable mise in the current shell session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deactivateCmd).Standalone()

	deactivateCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(deactivateCmd)
}
