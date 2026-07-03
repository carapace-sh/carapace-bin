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
	rootCmd.Flags().BoolS("D", "D", false, "delete a history event from the internal history list")
	rootCmd.Flags().BoolS("E", "E", false, "print full time-date stamps in European format")
	rootCmd.Flags().BoolS("I", "I", false, "restrict to only internal events")
	rootCmd.Flags().BoolS("L", "L", false, "restrict to only local events")
	rootCmd.Flags().BoolS("P", "P", false, "pop the saved history off the stack")
	rootCmd.Flags().BoolS("R", "R", false, "read history from a file")
	rootCmd.Flags().BoolS("W", "W", false, "write current history to a file")
	rootCmd.Flags().BoolS("d", "d", false, "print timestamps for each event")
	rootCmd.Flags().StringS("e", "e", "", "specify editor to use")
	rootCmd.Flags().BoolS("f", "f", false, "print full time-date stamps in US format")
	rootCmd.Flags().BoolS("i", "i", false, "print full time-date stamps in ISO8601 format")
	rootCmd.Flags().BoolS("l", "l", false, "list commands instead of editing")
	rootCmd.Flags().StringS("m", "m", "", "take the first argument as a pattern and only matching events are considered")
	rootCmd.Flags().BoolS("n", "n", false, "omit command numbers from display")
	rootCmd.Flags().BoolS("p", "p", false, "push the current history onto the stack")
	rootCmd.Flags().BoolS("r", "r", false, "reverse the order of the events")
	rootCmd.Flags().BoolS("s", "s", false, "equivalent to -e -")
	rootCmd.Flags().StringS("t", "t", "", "print time-date stamps in the given format")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"e": carapace.ActionExecutables(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Usage("first event number or string"),
	)
}
