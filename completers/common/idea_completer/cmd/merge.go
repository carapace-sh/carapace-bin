package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Open the Merge dialog to merge the specified files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mergeCmd).Standalone()

	rootCmd.AddCommand(mergeCmd)

	carapace.Gen(mergeCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
