package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eval",
	Short: "Evaluate code as string",
	Long:  "https://elv.sh/ref/builtin.html#eval",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().StringS("ns", "ns", "", "namespace to evaluate in")
	rootCmd.Flags().StringS("on-end", "on-end", "", "callback on evaluation end")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Usage("code to evaluate"),
	)
}
