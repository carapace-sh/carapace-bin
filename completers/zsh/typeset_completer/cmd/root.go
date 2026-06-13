package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "typeset",
	Short: "Define variables and set attributes",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-typeset",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("E", "E", false, "use scientific notation")
	rootCmd.Flags().BoolS("L", "L", false, "left justify")
	rootCmd.Flags().BoolS("R", "R", false, "right justify")
	rootCmd.Flags().BoolS("Z", "Z", false, "zero fill")
	rootCmd.Flags().BoolS("f", "f", false, "refer to functions")
	rootCmd.Flags().BoolS("h", "h", false, "use hash notation")
	rootCmd.Flags().BoolS("i", "i", false, "use integer arithmetic")
	rootCmd.Flags().BoolS("l", "l", false, "lowercase conversion")
	rootCmd.Flags().BoolS("r", "r", false, "make readonly")
	rootCmd.Flags().BoolS("t", "t", false, "mark as tagged")
	rootCmd.Flags().BoolS("u", "u", false, "uppercase conversion")
	rootCmd.Flags().BoolS("x", "x", false, "export")
}
