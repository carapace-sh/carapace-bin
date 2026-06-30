package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dCmd = &cobra.Command{
	Use:   "d",
	Short: "Delete files from archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dCmd).Standalone()

	addCommonFlags(dCmd)
	addCommonFlagCompletions(dCmd)
	rootCmd.AddCommand(dCmd)

	carapace.Gen(dCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(dCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
