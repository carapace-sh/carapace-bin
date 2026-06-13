package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fish_indent",
	Short: "Indent fish code",
	Long:  "https://fishshell.com/docs/current/cmds/fish_indent.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("ansi", false, "colorize with ANSI")
	rootCmd.Flags().BoolS("c", "c", false, "check formatting")
	rootCmd.Flags().Bool("check", false, "check formatting")
	rootCmd.Flags().Bool("dump-parse-tree", false, "dump parse tree")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().Bool("html", false, "output HTML")
	rootCmd.Flags().BoolS("i", "i", false, "only reformat")
	rootCmd.Flags().Bool("no-indent", false, "only reformat")
	rootCmd.Flags().Bool("only-indent", false, "only indent")
	rootCmd.Flags().Bool("only-unindent", false, "only unindent")
	rootCmd.Flags().BoolS("v", "v", false, "display version")
	rootCmd.Flags().BoolS("w", "w", false, "indent in-place")
	rootCmd.Flags().Bool("write", false, "indent in-place")
}
