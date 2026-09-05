package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var deleteBufferCmd = &cobra.Command{
	Use:     "delete-buffer",
	Aliases: []string{"deleteb"},
	Short:   "delete a paste buffer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteBufferCmd).Standalone()

	deleteBufferCmd.Flags().StringS("b", "b", "", "specify target buffer name")
	rootCmd.AddCommand(deleteBufferCmd)

	carapace.Gen(deleteBufferCmd).FlagCompletion(carapace.ActionMap{
		"b": tmux.ActionBuffers(),
	})
}
