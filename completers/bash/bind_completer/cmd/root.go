package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bind",
	Short: "Set Readline key bindings and variables",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-bind",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("P", "P", false, "list function names and bindings")
	rootCmd.Flags().BoolS("S", "S", false, "list key sequences that invoke macros and their values")
	rootCmd.Flags().BoolS("V", "V", false, "list variable names and values")
	rootCmd.Flags().BoolS("X", "X", false, "list key sequences bound with -x and associated commands in a form that can be reused as input")
	rootCmd.Flags().StringS("f", "f", "", "read key bindings from filename")
	rootCmd.Flags().BoolS("l", "l", false, "list names of functions")
	rootCmd.Flags().StringS("m", "m", "", "use keymap as the keymap for the duration of this command")
	rootCmd.Flags().BoolS("p", "p", false, "list functions and bindings in a form that can be reused as input")
	rootCmd.Flags().StringS("q", "q", "", "query about which keys invoke the named function")
	rootCmd.Flags().StringS("r", "r", "", "remove the binding for keyseq")
	rootCmd.Flags().BoolS("s", "s", false, "list key sequences that invoke macros and their values in a form that can be reused as input")
	rootCmd.Flags().StringS("u", "u", "", "unbind all keys which are bound to the named function")
	rootCmd.Flags().BoolS("v", "v", false, "list variable names and values in a form that can be reused as input")
	rootCmd.Flags().StringS("x", "x", "", "cause shell-command to be executed when keyseq is entered")

	// TODO function completion
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionFiles(),
		"m": carapace.ActionValuesDescribed(
			"emacs", "emacs keymap",
			"emacs-standard", "emacs-standard keymap",
			"emacs-meta", "emacs-meta keymap",
			"emacs-ctlx", "emacs-ctlx keymap",
			"vi", "vi keymap",
			"vi-move", "vi-move keymap",
			"vi-command", "vi-command keymap",
			"vi-insert", "vi-insert keymap",
		),
	})
}
