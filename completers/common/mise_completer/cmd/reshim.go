package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reshimCmd = &cobra.Command{
	Use:   "reshim",
	Short: "Create/recreate shims",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reshimCmd).Standalone()

	reshimCmd.Flags().BoolP("force", "f", false, "Force regeneration of all shims")
	rootCmd.AddCommand(reshimCmd)
}
