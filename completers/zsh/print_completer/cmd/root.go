package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "print",
	Short: "Display a line of text",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-print",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("C", "C", false, "print the arguments in columns")
	rootCmd.Flags().BoolS("D", "D", false, "treat the arguments as paths, replacing directory prefixes with ~ expressions")
	rootCmd.Flags().BoolS("N", "N", false, "print the arguments separated and terminated by nulls")
	rootCmd.Flags().BoolS("O", "O", false, "print the arguments sorted in descending order")
	rootCmd.Flags().BoolS("P", "P", false, "perform prompt expansion")
	rootCmd.Flags().BoolS("R", "R", false, "emulate the BSD echo command")
	rootCmd.Flags().BoolS("S", "S", false, "place the results in the history list as a single command line")
	rootCmd.Flags().BoolS("X", "X", false, "expand all tabs in the printed string")
	rootCmd.Flags().BoolS("a", "a", false, "print arguments with the column incrementing first")
	rootCmd.Flags().BoolS("b", "b", false, "recognize all the escape sequences defined for the bindkey command")
	rootCmd.Flags().BoolS("c", "c", false, "print the arguments in columns")
	rootCmd.Flags().BoolS("e", "e", false, "process escape sequences after -R")
	rootCmd.Flags().StringS("f", "f", "", "print the arguments as described by printf")
	rootCmd.Flags().BoolS("i", "i", false, "sort case-independently when used with -o or -O")
	rootCmd.Flags().BoolS("l", "l", false, "print the arguments separated by newlines instead of spaces")
	rootCmd.Flags().BoolS("m", "m", false, "take the first argument as a pattern and remove non-matching arguments")
	rootCmd.Flags().BoolS("n", "n", false, "do not add a newline to the output")
	rootCmd.Flags().BoolS("o", "o", false, "print the arguments sorted in ascending order")
	rootCmd.Flags().BoolS("p", "p", false, "print the arguments to the input of the coprocess")
	rootCmd.Flags().BoolS("r", "r", false, "ignore the escape conventions of echo")
	rootCmd.Flags().BoolS("s", "s", false, "place the results in the history list instead of on the standard output")
	rootCmd.Flags().StringS("u", "u", "", "print the arguments to file descriptor n")
	rootCmd.Flags().StringS("v", "v", "", "store the printed arguments as the value of the parameter name")
	rootCmd.Flags().StringS("x", "x", "", "expand leading tabs on each line of output")
	rootCmd.Flags().BoolS("z", "z", false, "push the arguments onto the editing buffer stack")
}
