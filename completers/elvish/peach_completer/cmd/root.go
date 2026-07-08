package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "peach",
	Short: "Apply a function to inputs in parallel",
	Long:  "https://elv.sh/ref/builtin.html#peach",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("num-workers", "num-workers", "", "number of parallel workers")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Usage("function to apply"),
	)
}
