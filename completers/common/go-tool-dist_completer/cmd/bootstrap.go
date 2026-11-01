package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "rebuild everything",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrapCmd).Standalone()

	bootstrapCmd.Flags().BoolS("a", "a", false, "rebuild all")
	bootstrapCmd.Flags().BoolS("d", "d", false, "enable debugging of bootstrap process")
	bootstrapCmd.Flags().BoolS("distpack", "distpack", false, "write distribution files to pkg/distpack")
	bootstrapCmd.Flags().BoolS("force", "force", false, "build even if the port is marked as broken")
	bootstrapCmd.Flags().BoolS("no-banner", "no-banner", false, "do not print banner")
	bootstrapCmd.Flags().BoolS("no-clean", "no-clean", false, "print deprecation warning")
	bootstrapCmd.Flags().BoolS("v", "v", false, "verbosity")
	rootCmd.AddCommand(bootstrapCmd)
}
