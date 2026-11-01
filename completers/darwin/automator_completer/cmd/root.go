package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "automator",
	Short: "run Automator workflow",
	Long:  "https://keith.github.io/xcode-manpages/automator.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("D", "D", "", "Set variable name=value")
	rootCmd.Flags().StringS("i", "i", "", "Set input as the input to workflow")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
