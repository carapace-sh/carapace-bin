package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unpairCmd = &cobra.Command{
	Use:   "unpair",
	Short: "Unpair a watch and phone pair",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unpairCmd).Standalone()
	rootCmd.AddCommand(unpairCmd)
	carapace.Gen(unpairCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
