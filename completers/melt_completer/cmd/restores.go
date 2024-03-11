package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Recreate a key using the given seed phrase",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restoreCmd).Standalone()

	restoreCmd.Flags().BoolP("help", "h", false, "help for restore")
	restoreCmd.Flags().StringP("seed", "s", "", "Seed phrase (default \"-\")")
	rootCmd.AddCommand(restoreCmd)

	carapace.Gen(restoreCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
