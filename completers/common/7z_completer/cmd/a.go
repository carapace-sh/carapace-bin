package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aCmd = &cobra.Command{
	Use:   "a",
	Short: "Add files to archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aCmd).Standalone()

	addCommonFlags(aCmd)
	addCommonFlagCompletions(aCmd)
	rootCmd.AddCommand(aCmd)

	carapace.Gen(aCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(aCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
