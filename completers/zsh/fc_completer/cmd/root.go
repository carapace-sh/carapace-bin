package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fc",
	Short: "Fix command - edit and re-execute commands from history",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-fc",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "append current history to history file")
	rootCmd.Flags().BoolS("D", "D", false, "don't run, just print")
	rootCmd.Flags().BoolS("W", "W", false, "append current history to history file")
	rootCmd.Flags().BoolS("e", "e", false, "specify editor to use")
	rootCmd.Flags().BoolS("l", "l", false, "list commands instead of editing")
	rootCmd.Flags().BoolS("n", "n", false, "omit command numbers from display")
	rootCmd.Flags().BoolS("r", "r", false, "reverse the order of commands")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Usage("first event number or string"),
	)
}
