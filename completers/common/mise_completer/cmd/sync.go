// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Use tool versions installed by other version managers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syncCmd).Standalone()

	syncCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(syncCmd)
}
