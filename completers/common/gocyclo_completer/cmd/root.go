package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocyclo",
	Short: "Calculate cyclomatic complexities of Go functions",
	Long:  "https://github.com/fzipp/gocyclo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("avg", "avg", false, "show the average complexity over all functions")
	rootCmd.Flags().BoolS("avg-short", "avg-short", false, "show the average complexity over all functions without a label")
	rootCmd.Flags().StringS("ignore", "ignore", "", "exclude files matching the given regular expression")
	rootCmd.Flags().IntS("over", "over", 0, "show functions with complexity > N only")
	rootCmd.Flags().IntS("top", "top", -1, "show the top N most complex functions only")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
