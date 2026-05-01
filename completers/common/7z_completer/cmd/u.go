package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uCmd = &cobra.Command{
	Use:   "u",
	Short: "Update files to archive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uCmd).Standalone()

	addCommonFlags(uCmd)
	addCommonFlagCompletions(uCmd)
	rootCmd.AddCommand(uCmd)

	carapace.Gen(uCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(uCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
