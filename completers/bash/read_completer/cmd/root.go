package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a line from the standard input and split it into fields",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-read",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("N", "N", "", "read returns after reading exactly nchars characters")
	rootCmd.Flags().StringS("d", "d", "", "continue until the first character of delim is read instead of newline")
	rootCmd.Flags().BoolS("e", "e", false, "use Readline to obtain the line in an interactive shell")
	rootCmd.Flags().StringS("i", "i", "", "text is placed in the initial edit buffer as if it had been entered")
	rootCmd.Flags().StringS("n", "n", "", "read returns after reading nchars characters")
	rootCmd.Flags().StringS("p", "p", "", "display prompt on standard error without a trailing newline before reading")
	rootCmd.Flags().BoolS("r", "r", false, "do not allow backslashes to escape any characters")
	rootCmd.Flags().BoolS("s", "s", false, "do not echo input coming from a terminal")
	rootCmd.Flags().StringS("t", "t", "", "time out and return failure if a complete line of input is not read within timeout seconds")
	rootCmd.Flags().StringS("u", "u", "", "read from file descriptor fd instead of standard input")
}
