package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "history",
	Short: "Show and manipulate command history",
	Long:  "https://fishshell.com/docs/current/cmds/history.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("C", "C", false, "case-sensitive search")
	rootCmd.Flags().Bool("case-sensitive", false, "case-sensitive search")
	rootCmd.Flags().BoolS("c", "c", false, "search by contains")
	rootCmd.Flags().Bool("contains", false, "search by contains")
	rootCmd.Flags().String("color", "", "color output")
	rootCmd.Flags().BoolS("e", "e", false, "exact match")
	rootCmd.Flags().Bool("exact", false, "exact match")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().StringS("n", "n", "", "limit results")
	rootCmd.Flags().String("max", "", "limit results")
	rootCmd.Flags().BoolS("p", "p", false, "search by prefix")
	rootCmd.Flags().Bool("prefix", false, "search by prefix")
	rootCmd.Flags().BoolS("R", "R", false, "oldest to newest")
	rootCmd.Flags().Bool("reverse", false, "oldest to newest")
	rootCmd.Flags().StringS("t", "t", "", "prepend timestamp")
	rootCmd.Flags().String("show-time", "", "prepend timestamp")
	rootCmd.Flags().BoolS("z", "z", false, "NUL-terminated output")
	rootCmd.Flags().Bool("null", false, "NUL-terminated output")
}
