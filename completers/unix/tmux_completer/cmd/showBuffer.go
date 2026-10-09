package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var showBufferCmd = &cobra.Command{
	Use:     "show-buffer",
	Aliases: []string{"showb"},
	Short:   "display the contents of a paste buffer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showBufferCmd).Standalone()

	showBufferCmd.Flags().StringS("b", "b", "", "specify target buffer name")
	rootCmd.AddCommand(showBufferCmd)

	carapace.Gen(showBufferCmd).FlagCompletion(carapace.ActionMap{
		"b": tmux.ActionBuffers(),
	})
}
