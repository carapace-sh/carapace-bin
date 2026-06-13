package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "printf",
	Short: "Display formatted output",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-printf",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("v", "v", false, "assign output to variable")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Usage("format string"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValues(),
	)
}
