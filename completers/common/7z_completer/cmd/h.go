package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hCmd = &cobra.Command{
	Use:   "h",
	Short: "Calculate hash values for files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hCmd).Standalone()

	addCommonFlags(hCmd)
	addCommonFlagCompletions(hCmd)
	rootCmd.AddCommand(hCmd)

	carapace.Gen(hCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
