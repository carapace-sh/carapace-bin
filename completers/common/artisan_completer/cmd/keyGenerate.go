package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyGenerateCmd = &cobra.Command{
	Use:   "key:generate [--show] [--force]",
	Short: "Set the application key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyGenerateCmd).Standalone()

	keyGenerateCmd.Flags().Bool("force", false, "Force the operation to run when in production")
	keyGenerateCmd.Flags().Bool("show", false, "Display the key instead of modifying files")
	rootCmd.AddCommand(keyGenerateCmd)
}
