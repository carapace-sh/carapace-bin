package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "range",
	Short: "Output a range of numbers",
	Long:  "https://elv.sh/ref/builtin.html#range",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("step", "step", "", "step size")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Usage("end or start"),
		carapace.ActionValues().Usage("end"),
	)
}
