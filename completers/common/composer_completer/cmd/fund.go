package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fundCmd = &cobra.Command{
	Use:   "fund",
	Short: "Discover how to help fund the maintenance of your dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fundCmd).Standalone()

	fundCmd.Flags().StringP("format", "f", "", "Format of the output: text or json")

	rootCmd.AddCommand(fundCmd)

	carapace.Gen(fundCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "text"),
	})
}
