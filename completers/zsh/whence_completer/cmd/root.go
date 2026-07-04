package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whence",
	Short: "Identify command types",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-whence",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("S", "S", false, "print intermediate symlink steps")
	rootCmd.Flags().BoolS("a", "a", false, "search for all occurrences of name throughout the command path")
	rootCmd.Flags().BoolS("c", "c", false, "print the results in a csh-like format")
	rootCmd.Flags().BoolS("f", "f", false, "display the contents of a shell function")
	rootCmd.Flags().BoolS("m", "m", false, "take arguments as patterns and display information for each matching entry")
	rootCmd.Flags().BoolS("p", "p", false, "do a path search for name even if it is an alias, reserved word, shell function or builtin")
	rootCmd.Flags().BoolS("s", "s", false, "print the symlink-free pathname as well")
	rootCmd.Flags().BoolS("v", "v", false, "produce a more verbose report")
	rootCmd.Flags().BoolS("w", "w", false, "describe how word would be interpreted")
	rootCmd.Flags().StringS("x", "x", "", "expand tabs when outputting shell functions using the -c option")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionExecutables().FilterArgs(),
	)
}
