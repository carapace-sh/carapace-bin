package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "comp",
	Short: "compare the contents of two files byte by byte",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/comp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("a", "a", false, "ASCII comparison")
	rootCmd.Flags().BoolP("d", "d", false, "decimal comparison")
	rootCmd.Flags().BoolP("n", "n", false, "display line numbers")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
