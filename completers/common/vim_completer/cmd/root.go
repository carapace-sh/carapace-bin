package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vim",
	Short: "Vi IMproved, a programmer's text editor",
	Long:  "https://linux.die.net/man/1/vi",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "Start in Arabic mode")
	rootCmd.Flags().BoolS("C", "C", false, "Compatible with Vi: 'compatible'")
	rootCmd.Flags().BoolS("D", "D", false, "Debugging mode")
	rootCmd.Flags().BoolS("E", "E", false, "Improved Ex mode")
	rootCmd.Flags().BoolS("H", "H", false, "Start in Hebrew mode")
	rootCmd.Flags().BoolS("L", "L", false, "Same as -r")
	rootCmd.Flags().BoolS("M", "M", false, "Modifications in text not allowed")
	rootCmd.Flags().BoolS("N", "N", false, "Not fully Vi compatible: 'nocompatible'")
	rootCmd.Flags().StringS("O", "O", "", "Like -o but split vertically")
	rootCmd.Flags().BoolS("R", "R", false, "Readonly mode")
	rootCmd.Flags().StringS("S", "S", "", "Source file <session> after loading the first file")
	rootCmd.Flags().StringS("T", "T", "", "Set terminal type to <terminal>")
	rootCmd.Flags().StringS("V", "V", "", "Be verbose")
	rootCmd.Flags().StringS("W", "W", "", "Write all typed commands to file <scriptout>")
	rootCmd.Flags().BoolS("Z", "Z", false, "Restricted mode")
	rootCmd.Flags().BoolS("b", "b", false, "Binary mode")
	rootCmd.Flags().StringS("c", "c", "", "Execute <command> after loading the first file")
	rootCmd.Flags().Bool("clean", false, "'nocompatible', Vim defaults, no plugins, no viminfo")
	rootCmd.Flags().String("cmd", "", "Execute <command> before loading any vimrc file")
	rootCmd.Flags().BoolS("d", "d", false, "Diff mode")
	rootCmd.Flags().BoolS("e", "e", false, "Ex mode")
	rootCmd.Flags().BoolP("help", "h", false, "Print Help (this message) and exit")
	rootCmd.Flags().StringS("i", "i", "", "Use <viminfo> instead of .viminfo")
	rootCmd.Flags().BoolS("l", "l", false, "Lisp mode")
	rootCmd.Flags().String("log", "", "Start logging to <file> early")
	rootCmd.Flags().BoolS("m", "m", false, "Modifications (writing files) not allowed")
	rootCmd.Flags().BoolS("n", "n", false, "No swap file, use memory only")
	rootCmd.Flags().Bool("noplugin", false, "Do not load plugin scripts")
	rootCmd.Flags().Bool("not-a-term", false, "Skip warning for input/output not being a terminal")
	rootCmd.Flags().StringS("o", "o", "", "Open N windows")
	rootCmd.Flags().StringS("p", "p", "", "Open N tab pages")
	rootCmd.Flags().BoolS("r", "r", false, "List swap files or recover crashed session")
	rootCmd.Flags().StringS("s", "s", "", "Read Normal mode commands from file <scriptin>")
	rootCmd.Flags().String("startuptime", "", "Write startup timing messages to <file>")
	rootCmd.Flags().Bool("ttyfail", false, "Exit if input or output is not a terminal")
	rootCmd.Flags().StringS("u", "u", "", "Use <vimrc> instead of any .vimrc")
	rootCmd.Flags().BoolS("v", "v", false, "Vi mode")
	rootCmd.Flags().Bool("version", false, "Print version information and exit")
	rootCmd.Flags().StringS("w", "w", "", "Append all typed commands to file <scriptout>")
	rootCmd.Flags().BoolS("x", "x", false, "Edit encrypted files")
	rootCmd.Flags().BoolS("y", "y", false, "Easy mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"S":           carapace.ActionFiles(),
		"W":           carapace.ActionFiles(),
		"i":           carapace.ActionFiles(),
		"log":         carapace.ActionFiles(),
		"s":           carapace.ActionFiles(),
		"startuptime": carapace.ActionFiles(),
		"u":           carapace.ActionFiles(),
		"w":           carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(rootCmd).DashAnyCompletion(
		carapace.ActionPositional(rootCmd),
	)

}
