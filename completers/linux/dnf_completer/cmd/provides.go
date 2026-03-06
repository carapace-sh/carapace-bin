package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var providesCmd = &cobra.Command{
	Use:   "provides <package-spec>...",
	Short: "find what package provides the given value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providesCmd).Standalone()

	rootCmd.AddCommand(providesCmd)

	carapace.Gen(providesCmd).PositionalCompletion(
		carapace.ActionValues(), // TODO: complete with provides
	)
}
