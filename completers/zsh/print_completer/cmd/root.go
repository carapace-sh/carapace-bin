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

	rootCmd.Flags().BoolS("D", "D", false, "interpret escape sequences")
	rootCmd.Flags().BoolS("R", "R", false, "like -e")
	rootCmd.Flags().BoolS("n", "n", false, "do not output a trailing newline")
	rootCmd.Flags().BoolS("z", "z", false, "push arguments onto editing buffer stack")
}
