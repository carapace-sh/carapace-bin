package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var nextLayoutCmd = &cobra.Command{
	Use:     "next-layout",
	Aliases: []string{"nextl"},
	Short:   "move a window to the next layout",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nextLayoutCmd).Standalone()

	nextLayoutCmd.Flags().StringS("t", "t", "", "specify target window")
	rootCmd.AddCommand(nextLayoutCmd)

	carapace.Gen(nextLayoutCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionWindows(),
	})
}
