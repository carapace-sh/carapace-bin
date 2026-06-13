package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "set",
	Short: "Set or unset values of shell options and positional parameters",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-set",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "array assignment")
	rootCmd.Flags().BoolS("e", "e", false, "exit immediately on error")
	rootCmd.Flags().BoolS("f", "f", false, "disable file name generation")
	rootCmd.Flags().BoolS("h", "h", false, "remember command locations")
	rootCmd.Flags().BoolS("k", "k", false, "assign arguments as environment variables")
	rootCmd.Flags().BoolS("m", "m", false, "enable job control")
	rootCmd.Flags().BoolS("n", "n", false, "read but do not execute")
	rootCmd.Flags().BoolS("t", "t", false, "exit after one command")
	rootCmd.Flags().BoolS("u", "u", false, "treat unset variables as error")
	rootCmd.Flags().BoolS("v", "v", false, "print shell input lines")
	rootCmd.Flags().BoolS("x", "x", false, "print commands and arguments")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"A": carapace.ActionValues("name", "array"),
	})
}
