// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_unapplyCmd = &cobra.Command{
	Use:   "unapply",
	Short: "Remove what a config environment's bootstrap sections created",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_unapplyCmd).Standalone()

	bootstrap_unapplyCmd.Flags().BoolP("dry-run", "n", false, "Show what would be removed without removing anything")
	bootstrap_unapplyCmd.Flags().BoolP("force", "f", false, "Remove targets that changed since they were applied")
	bootstrap_unapplyCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_unapplyCmd.Flags().Bool("prompt-secrets", false, "Prompt securely for missing bootstrap secret inputs")
	bootstrap_unapplyCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt")
	bootstrapCmd.AddCommand(bootstrap_unapplyCmd)
}
