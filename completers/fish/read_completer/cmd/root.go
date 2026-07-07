package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "read",
	Short: "Read input into variables",
	Long:  "https://fishshell.com/docs/current/cmds/read.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("L", "L", false, "read into successive vars")
	rootCmd.Flags().StringS("P", "P", "", "literal prompt string")
	rootCmd.Flags().StringS("R", "R", "", "right prompt command")
	rootCmd.Flags().BoolS("S", "S", false, "shell highlighting")
	rootCmd.Flags().BoolS("U", "U", false, "universal variable")
	rootCmd.Flags().BoolS("a", "a", false, "store as list")
	rootCmd.Flags().Bool("array", false, "store as list (backward-compatible alias for --list)")
	rootCmd.Flags().StringS("c", "c", "", "initial string")
	rootCmd.Flags().String("command", "", "initial string")
	rootCmd.Flags().StringS("d", "d", "", "split on delimiter")
	rootCmd.Flags().String("delimiter", "", "split on delimiter")
	rootCmd.Flags().Bool("export", false, "export")
	rootCmd.Flags().BoolS("f", "f", false, "function-scoped variable")
	rootCmd.Flags().Bool("function", false, "function-scoped variable")
	rootCmd.Flags().BoolS("g", "g", false, "global variable")
	rootCmd.Flags().Bool("global", false, "global variable")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("l", "l", false, "local variable")
	rootCmd.Flags().Bool("line", false, "read into successive vars")
	rootCmd.Flags().Bool("list", false, "store as list")
	rootCmd.Flags().Bool("local", false, "local variable")
	rootCmd.Flags().StringS("n", "n", "", "return after N characters")
	rootCmd.Flags().String("nchars", "", "return after N characters")
	rootCmd.Flags().Bool("null", false, "NUL delimiter")
	rootCmd.Flags().StringS("p", "p", "", "custom prompt command")
	rootCmd.Flags().String("prompt", "", "custom prompt command")
	rootCmd.Flags().String("prompt-str", "", "literal prompt string")
	rootCmd.Flags().String("right-prompt", "", "right prompt command")
	rootCmd.Flags().BoolS("s", "s", false, "mask input")
	rootCmd.Flags().Bool("shell", false, "shell highlighting")
	rootCmd.Flags().Bool("silent", false, "mask input")
	rootCmd.Flags().BoolS("t", "t", false, "shell tokenization")
	rootCmd.Flags().Bool("tokenize", false, "shell tokenization")
	rootCmd.Flags().Bool("tokenize-raw", false, "shell tokenization without quote removal")
	rootCmd.Flags().BoolS("u", "u", false, "do not export")
	rootCmd.Flags().Bool("unexport", false, "do not export")
	rootCmd.Flags().Bool("universal", false, "universal variable")
	rootCmd.Flags().BoolS("x", "x", false, "export")
	rootCmd.Flags().BoolS("z", "z", false, "NUL delimiter")
}
