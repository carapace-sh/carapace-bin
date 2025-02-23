package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Print full output of recorded session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catCmd).Standalone()

	rootCmd.AddCommand(catCmd)

	carapace.Gen(catCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
