package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xCmd = &cobra.Command{
	Use:   "x",
	Short: "Extract files with full paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xCmd).Standalone()

	addCommonFlags(xCmd)
	addCommonFlagCompletions(xCmd)
	rootCmd.AddCommand(xCmd)

	carapace.Gen(xCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(xCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
