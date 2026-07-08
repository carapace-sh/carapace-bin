package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "order",
	Short: "Sort values",
	Long:  "https://elv.sh/ref/builtin.html#order",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("key", "key", "", "key function to apply to each element")
	rootCmd.Flags().StringS("less-than", "less-than", "", "custom less-than comparator")
	rootCmd.Flags().BoolS("reverse", "reverse", false, "reverse the order")
	rootCmd.Flags().BoolS("total", "total", false, "make the comparison total")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValues().Usage("values to sort"),
	)
}
