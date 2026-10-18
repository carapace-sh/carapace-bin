package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var renameWindowCmd = &cobra.Command{
	Use:     "rename-window",
	Aliases: []string{"renamew"},
	Short:   "rename a window",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(renameWindowCmd).Standalone()

	renameWindowCmd.Flags().StringS("t", "t", "", "specify target window")
	rootCmd.AddCommand(renameWindowCmd)

	carapace.Gen(renameWindowCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionWindows(),
	})
}
