package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ps",
	Short: "process status",
	Long:  "https://keith.github.io/xcode-manpages/ps.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "Select all processes")
	rootCmd.Flags().BoolS("C", "C", false, "Show command name")
	rootCmd.Flags().BoolS("E", "E", false, "Show environment")
	rootCmd.Flags().BoolS("L", "L", false, "List all thread formats")
	rootCmd.Flags().BoolS("M", "M", false, "Show memory usage")
	rootCmd.Flags().BoolS("S", "S", false, "Show swap usage")
	rootCmd.Flags().BoolS("T", "T", false, "Show thread count")
	rootCmd.Flags().BoolS("X", "X", false, "Do not show daemons")
	rootCmd.Flags().BoolS("a", "a", false, "Select all processes except session leaders")
	rootCmd.Flags().BoolS("e", "e", false, "Select all processes (equiv to -A)")
	rootCmd.Flags().BoolS("f", "f", false, "Full format listing")
	rootCmd.Flags().BoolS("h", "h", false, "Show hierarchical tree")
	rootCmd.Flags().BoolS("j", "j", false, "Show job control format")
	rootCmd.Flags().BoolS("l", "l", false, "Long format")
	rootCmd.Flags().BoolS("m", "m", false, "Sort by memory usage")
	rootCmd.Flags().BoolS("r", "r", false, "Sort by CPU usage")
	rootCmd.Flags().BoolS("v", "v", false, "Show virtual memory")
	rootCmd.Flags().BoolS("w", "w", false, "Wide output")
	rootCmd.Flags().BoolS("x", "x", false, "Include daemons")

	rootCmd.Flags().StringS("G", "G", "", "Select by real group ID")
	rootCmd.Flags().StringS("O", "O", "", "Add columns after pid")
	rootCmd.Flags().StringS("U", "U", "", "Select by real user ID")
	rootCmd.Flags().StringS("g", "g", "", "Select by session leader group")
	rootCmd.Flags().StringS("o", "o", "", "Specify output format")
	rootCmd.Flags().StringS("p", "p", "", "Select by PID")
	rootCmd.Flags().StringS("t", "t", "", "Select by tty")
	rootCmd.Flags().StringS("u", "u", "", "Select by effective user ID")
}
