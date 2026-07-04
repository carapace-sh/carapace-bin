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

	rootCmd.Flags().BoolS("A", "A", false, "the first name is taken as the name of an array and all words are assigned to it")
	rootCmd.Flags().BoolS("E", "E", false, "the input read is printed to the standard output")
	rootCmd.Flags().BoolS("c", "c", false, "read the words of the current command (completion only)")
	rootCmd.Flags().StringS("d", "d", "", "input is terminated by the first character of delim instead of by newline")
	rootCmd.Flags().BoolS("e", "e", false, "the input read is printed to the standard output but no input is assigned")
	rootCmd.Flags().StringS("k", "k", "", "read only one (or num) characters")
	rootCmd.Flags().BoolS("l", "l", false, "the whole line is assigned as a scalar (completion only)")
	rootCmd.Flags().BoolS("n", "n", false, "together with -c, the number of the word the cursor is on is read")
	rootCmd.Flags().BoolS("p", "p", false, "input is read from the coprocess")
	rootCmd.Flags().BoolS("q", "q", false, "read only one character from the terminal and set name to y or n")
	rootCmd.Flags().BoolS("r", "r", false, "raw mode: backslashes do not quote the following character")
	rootCmd.Flags().BoolS("s", "s", false, "don't echo back characters if reading from the terminal")
	rootCmd.Flags().BoolS("t", "t", false, "test if input is available before attempting to read")
	rootCmd.Flags().StringS("u", "u", "", "input is read from file descriptor n")
	rootCmd.Flags().BoolS("z", "z", false, "read one entry from the editor buffer stack")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Usage("variable name"),
	)
}
