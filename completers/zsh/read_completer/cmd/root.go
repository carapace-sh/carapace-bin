package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a line from the standard input",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-read",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("0", "0", false, "delimiter is NUL")
	rootCmd.Flags().BoolS("A", "A", false, "array")
	rootCmd.Flags().BoolS("r", "r", false, "raw mode")
	rootCmd.Flags().BoolS("s", "s", false, "no output")
	rootCmd.Flags().BoolS("t", "t", false, "timeout")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Usage("variable name"),
	)
}
