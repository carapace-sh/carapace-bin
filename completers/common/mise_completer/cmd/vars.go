// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var varsCmd = &cobra.Command{
	Use:   "vars",
	Short: "List, set, and ask for config vars",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(varsCmd).Standalone()

	varsCmd.Flags().BoolP("help", "h", false, "Print help")
	varsCmd.Flags().Bool("no-header", false, "Do not print the table header")
	rootCmd.AddCommand(varsCmd)
}
