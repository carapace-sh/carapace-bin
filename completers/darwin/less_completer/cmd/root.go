package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "less",
	Short: "opposite of more",
	Long:  "https://man.freebsd.org/cgi/man.cgi?less",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "Help")
	rootCmd.Flags().BoolS("A", "A", false, "Search starts after last line")
	rootCmd.Flags().BoolS("B", "B", false, "Auto buffers")
	rootCmd.Flags().BoolS("C", "C", false, "Clear screen from top")
	rootCmd.Flags().StringS("D", "D", "", "Color")
	rootCmd.Flags().BoolS("E", "E", false, "Quit at EOF (including less) ")
	rootCmd.Flags().BoolS("F", "F", false, "Quit if one screen")
	rootCmd.Flags().BoolS("G", "G", false, "Do not highlight search matches")
	rootCmd.Flags().BoolS("I", "I", false, "Ignore case in searches (all)")
	rootCmd.Flags().BoolS("J", "J", false, "Status column")
	rootCmd.Flags().BoolS("K", "K", false, "Quit on interrupt")
	rootCmd.Flags().BoolS("L", "L", false, "No lessopen")
	rootCmd.Flags().BoolS("M", "M", false, "Very long prompt")
	rootCmd.Flags().BoolS("N", "N", false, "Line numbers (always)")
	rootCmd.Flags().StringS("O", "O", "", "Log file (overwrite)")
	rootCmd.Flags().StringS("P", "P", "", "Prompt")
	rootCmd.Flags().BoolS("Q", "Q", false, "Completely quiet")
	rootCmd.Flags().BoolS("R", "R", false, "Raw control chars (color)")
	rootCmd.Flags().BoolS("S", "S", false, "Chop long lines")
	rootCmd.Flags().StringS("T", "T", "", "Tags file")
	rootCmd.Flags().BoolS("U", "U", false, "Underline special (all)")
	rootCmd.Flags().BoolS("V", "V", false, "Version")
	rootCmd.Flags().BoolS("W", "W", false, "Highlight first line after any forward movement")
	rootCmd.Flags().BoolS("X", "X", false, "No termcap init/deinit")
	rootCmd.Flags().StringS("Y", "Y", "", "Max forward scroll (same as -y)")
	rootCmd.Flags().BoolS("a", "a", false, "Search skips screen")
	rootCmd.Flags().StringS("b", "b", "", "Number of buffers")
	rootCmd.Flags().BoolS("c", "c", false, "Clear screen")
	rootCmd.Flags().BoolS("d", "d", false, "Dumb terminal")
	rootCmd.Flags().BoolS("e", "e", false, "Quit at EOF")
	rootCmd.Flags().BoolS("f", "f", false, "Force open non-regular files")
	rootCmd.Flags().BoolS("g", "g", false, "Highlight search matches")
	rootCmd.Flags().StringS("h", "h", "", "Maximum back scroll")
	rootCmd.Flags().BoolS("i", "i", false, "Ignore case in searches")
	rootCmd.Flags().StringS("j", "j", "", "Target line")
	rootCmd.Flags().StringS("k", "k", "", "Lesskey file")
	rootCmd.Flags().BoolS("m", "m", false, "Long prompt")
	rootCmd.Flags().BoolS("n", "n", false, "Line numbers")
	rootCmd.Flags().StringS("o", "o", "", "Log file")
	rootCmd.Flags().StringS("p", "p", "", "Search pattern")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet")
	rootCmd.Flags().BoolS("r", "r", false, "Raw control chars")
	rootCmd.Flags().BoolS("s", "s", false, "Squeeze blank lines")
	rootCmd.Flags().StringS("t", "t", "", "Tag")
	rootCmd.Flags().BoolS("u", "u", false, "Underline special")
	rootCmd.Flags().BoolS("w", "w", false, "Highlight first line after move")
	rootCmd.Flags().StringS("x", "x", "", "Tab stops")
	rootCmd.Flags().StringS("y", "y", "", "Max forward scroll")
	rootCmd.Flags().StringS("z", "z", "", "Window size")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"O": carapace.ActionFiles(),
		"k": carapace.ActionFiles(),
		"o": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
