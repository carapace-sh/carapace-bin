package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rnCmd = &cobra.Command{
	Use:   "rn",
	Short: "Rename files in archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rnCmd).Standalone()

	addCommonFlags(rnCmd)
	addCommonFlagCompletions(rnCmd)
	rootCmd.AddCommand(rnCmd)

	carapace.Gen(rnCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
