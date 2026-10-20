// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Show what applying bootstrap resources would change",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_planCmd).Standalone()

	bootstrap_planCmd.Flags().Bool("detailed-exitcode", false, "Exit 0 if nothing would change, 2 if something would, 1 on errors or unknown state")
	bootstrap_planCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_planCmd.Flags().BoolP("json", "J", false, "Output a stable machine-readable plan in JSON format")
	bootstrap_planCmd.Flags().Bool("prompt-secrets", false, "Prompt securely for missing bootstrap secret inputs")
	bootstrapCmd.AddCommand(bootstrap_planCmd)
}
