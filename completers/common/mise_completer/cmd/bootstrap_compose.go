// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_composeCmd = &cobra.Command{
	Use:   "compose",
	Short: "Manage Docker Compose projects from `[bootstrap.compose]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_composeCmd).Standalone()

	bootstrap_composeCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_composeCmd)
}
