package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tCmd = &cobra.Command{
	Use:   "t",
	Short: "Test integrity of archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tCmd).Standalone()

	addCommonFlags(tCmd)
	addCommonFlagCompletions(tCmd)
	rootCmd.AddCommand(tCmd)

	carapace.Gen(tCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
