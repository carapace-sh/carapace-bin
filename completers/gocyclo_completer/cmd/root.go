package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocyclo",
	Short: "Calculate cyclomatic complexities of Go functions",
	Long:  "https://github.com/fzipp/gocyclo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("avg", false, "show the average complexity over all functions")
	rootCmd.Flags().Bool("avg-short", false, "show the average complexity over all functions without a label")
	rootCmd.Flags().String("ignore", "", "exclude files matching the given regular expression")
	rootCmd.Flags().String("over", "", "show functions with complexity > N only")
	rootCmd.Flags().String("top", "", "show the top N most complex functions only")
	rootCmd.Flags().Bool("total", false, "show the total complexity for all functions")
	rootCmd.Flags().Bool("total-short", false, "show the total complexity for all functions without a label")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
