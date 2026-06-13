package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "echo",
	Short: "Display a line of text",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-echo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("E", "E", false, "disable interpretation of backslash escapes")
	rootCmd.Flags().BoolS("e", "e", false, "enable interpretation of backslash escapes")
	rootCmd.Flags().BoolS("n", "n", false, "do not output a trailing newline")
}
