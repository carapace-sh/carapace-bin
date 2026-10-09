package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var previousLayoutCmd = &cobra.Command{
	Use:     "previous-layout",
	Aliases: []string{"prevl"},
	Short:   "move a window to the previous layout",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(previousLayoutCmd).Standalone()

	previousLayoutCmd.Flags().StringS("t", "t", "", "specify target window")
	rootCmd.AddCommand(previousLayoutCmd)

	carapace.Gen(previousLayoutCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionWindows(),
	})
}
