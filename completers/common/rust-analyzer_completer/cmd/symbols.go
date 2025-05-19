package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var symbolsCmd = &cobra.Command{
	Use:   "symbols",
	Short: "Parse stdin and print the list of symbols",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(symbolsCmd).Standalone()

	rootCmd.AddCommand(symbolsCmd)
}
