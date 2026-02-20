package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Open the diff viewer to see the differences between two specified files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
