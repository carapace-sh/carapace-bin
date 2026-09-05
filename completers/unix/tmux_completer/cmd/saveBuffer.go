package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var saveBufferCmd = &cobra.Command{
	Use:     "save-buffer",
	Aliases: []string{"saveb"},
	Short:   "save a paste buffer to a file",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(saveBufferCmd).Standalone()

	saveBufferCmd.Flags().BoolS("a", "a", false, "append to rather than overwriting file")
	saveBufferCmd.Flags().StringS("b", "b", "", "specify a target buffer index")
	rootCmd.AddCommand(saveBufferCmd)

	carapace.Gen(saveBufferCmd).FlagCompletion(carapace.ActionMap{
		"b": tmux.ActionBuffers(),
	})

	carapace.Gen(saveBufferCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
