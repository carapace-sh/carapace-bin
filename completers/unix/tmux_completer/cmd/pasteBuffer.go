package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var pasteBufferCmd = &cobra.Command{
	Use:     "paste-buffer",
	Aliases: []string{"pasteb"},
	Short:   "insert a paste buffer into the window",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pasteBufferCmd).Standalone()

	pasteBufferCmd.Flags().BoolS("S", "S", false, "disable sanitization of control characters")
	pasteBufferCmd.Flags().StringS("b", "b", "", "specify buffer")
	pasteBufferCmd.Flags().BoolS("d", "d", false, "remove buffer from stack after pasting")
	pasteBufferCmd.Flags().BoolS("p", "p", false, "use bracketed paste mode if the application requested it")
	pasteBufferCmd.Flags().BoolS("r", "r", false, "don't replace LF with CR when pasting")
	pasteBufferCmd.Flags().StringS("s", "s", "", "specify separator")
	pasteBufferCmd.Flags().StringS("t", "t", "", "specify target pane")
	rootCmd.AddCommand(pasteBufferCmd)

	carapace.Gen(pasteBufferCmd).FlagCompletion(carapace.ActionMap{
		"b": tmux.ActionBuffers(),
		"t": tmux.ActionPanes(),
	})
}
