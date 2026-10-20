// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "[experimental] Inspect the project's secrets source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsCmd).Standalone()

	secretsCmd.Flags().BoolP("help", "h", false, "Print help")
	secretsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	secretsCmd.Flags().Bool("no-header", false, "Do not print the table header")
	rootCmd.AddCommand(secretsCmd)
}
