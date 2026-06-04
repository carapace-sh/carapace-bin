package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "readarray",
	Short: "Read lines from standard input into an indexed array variable",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-readarray",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("C", "C", "", "evaluate callback each time quantum lines are read")
	rootCmd.Flags().StringS("O", "O", "", "begin assigning to array at index origin")
	rootCmd.Flags().StringS("c", "c", "", "specify the number of lines read between each call to callback")
	rootCmd.Flags().StringS("d", "d", "", "read lines until the first character of delim is read instead of newline")
	rootCmd.Flags().StringS("n", "n", "", "copy at most count lines")
	rootCmd.Flags().BoolS("s", "s", false, "treat the first line read as the number of lines to skip")
	rootCmd.Flags().StringS("t", "t", "", "remove a trailing delim from each line read")
	rootCmd.Flags().StringS("u", "u", "", "read lines from file descriptor fd instead of the standard input")
}
