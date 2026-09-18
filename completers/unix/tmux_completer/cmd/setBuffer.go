package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var setBufferCmd = &cobra.Command{
	Use:     "set-buffer",
	Aliases: []string{"setb"},
	Short:   "set contents of a paste buffer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setBufferCmd).Standalone()

	setBufferCmd.Flags().BoolS("a", "a", false, "append to rather than overwriting target buffer")
	setBufferCmd.Flags().StringS("b", "b", "", "specify target buffer name")
	setBufferCmd.Flags().StringS("n", "n", "", "specify new buffer name")
	setBufferCmd.Flags().StringS("t", "t", "", "specify target client")
	setBufferCmd.Flags().BoolS("w", "w", false, "also send the buffer to the clipboard")
	rootCmd.AddCommand(setBufferCmd)

	carapace.Gen(setBufferCmd).FlagCompletion(carapace.ActionMap{
		"b": tmux.ActionBuffers(),
		"t": tmux.ActionClients(),
	})
}
