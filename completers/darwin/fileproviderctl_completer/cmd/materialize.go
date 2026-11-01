package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var materializeCmd = &cobra.Command{
	Use:   "materialize",
	Short: "cause the specified item to be written on disk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(materializeCmd).Standalone()
	rootCmd.AddCommand(materializeCmd)
	carapace.Gen(materializeCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
