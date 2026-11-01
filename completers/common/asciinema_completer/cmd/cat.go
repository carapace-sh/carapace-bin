package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Concatenate multiple recordings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catCmd).Standalone()

	catCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(catCmd)

	carapace.Gen(catCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
