package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var loadBufferCmd = &cobra.Command{
	Use:     "load-buffer",
	Aliases: []string{"loadb"},
	Short:   "load a file into a paste buffer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loadBufferCmd).Standalone()

	loadBufferCmd.Flags().StringS("b", "b", "", "specify target buffer name")
	loadBufferCmd.Flags().StringS("t", "t", "", "specify target client")
	loadBufferCmd.Flags().BoolS("w", "w", false, "also send the buffer to the clipboard")
	rootCmd.AddCommand(loadBufferCmd)

	carapace.Gen(loadBufferCmd).FlagCompletion(carapace.ActionMap{
		"b": tmux.ActionBuffers(),
		"t": tmux.ActionClients(),
	})

	carapace.Gen(loadBufferCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
