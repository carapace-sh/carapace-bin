package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "compare",
	Short: "Compare two values",
	Long:  "https://elv.sh/ref/builtin.html#compare",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("total", "total", false, "make the comparison total")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Usage("first value"),
		carapace.ActionValues().Usage("second value"),
	)
}
