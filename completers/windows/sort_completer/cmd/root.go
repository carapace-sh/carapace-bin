package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sort",
	Short: "read input, sort data, and write results to the screen or file",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/sort",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("r", "r", false, "reverse the sort order (Z to A, 9 to 0)")
	rootCmd.Flags().BoolP("unique", "u", false, "remove all but one of any duplicate lines")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
