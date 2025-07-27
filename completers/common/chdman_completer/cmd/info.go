package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "displays information about a CHD",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().StringP("input", "i", "", "input file name")
	infoCmd.Flags().BoolP("verbose", "v", false, "output additional information")
	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).FlagCompletion(carapace.ActionMap{
		"input": carapace.ActionFiles(),
	})
}
