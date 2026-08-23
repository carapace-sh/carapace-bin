package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(restoreCmd)
}

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "restore an item from a backup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restoreCmd).Standalone()

	restoreCmd.Flags().BoolP("verbose", "v", false, "Verbose output")

	carapace.Gen(restoreCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}