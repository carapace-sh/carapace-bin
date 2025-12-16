package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "declare",
	Short: "Declare variables and give them attributes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "each name is an associative array variable")
	rootCmd.Flags().BoolS("F", "F", false, "inhibit  the display of function definitions")
	rootCmd.Flags().BoolS("I", "I", false, "inherit the attributes and value of any existing variable with the same name")
	rootCmd.Flags().BoolS("a", "a", false, "each name is an indexed array variable")
	rootCmd.Flags().BoolS("f", "f", false, "each name refers to a shell function")
	rootCmd.Flags().BoolS("g", "g", false, "forces variables to be created or modified at the global scope")
	rootCmd.Flags().BoolS("i", "i", false, "the variable is to be treated as an integer")
	rootCmd.Flags().BoolS("l", "l", false, "all upper-case characters are converted to lower-case")
	rootCmd.Flags().BoolS("n", "n", false, "give each name the nameref attribute")
	rootCmd.Flags().BoolS("p", "p", false, "option will display the attributes and values of each name")
	rootCmd.Flags().BoolS("r", "r", false, "make names readonly")
	rootCmd.Flags().BoolS("t", "t", false, "give each name the trace attribute")
	rootCmd.Flags().BoolS("u", "u", false, "all lower-case characters are converted to upper-case")
	rootCmd.Flags().BoolS("x", "x", false, "mark each name for export to subsequent commands via the environment")

	// TODO shell.Variables
}
