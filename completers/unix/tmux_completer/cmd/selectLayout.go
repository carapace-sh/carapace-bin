package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var selectLayoutCmd = &cobra.Command{
	Use:     "select-layout",
	Aliases: []string{"selectl"},
	Short:   "choose a layout for a pane",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selectLayoutCmd).Standalone()

	selectLayoutCmd.Flags().BoolS("E", "E", false, "spread the current pane and any panes next to it out evenly")
	selectLayoutCmd.Flags().BoolS("n", "n", false, "behave like next-layout")
	selectLayoutCmd.Flags().BoolS("o", "o", false, "revert to previous layout")
	selectLayoutCmd.Flags().BoolS("p", "p", false, "behave like previous-layout")
	selectLayoutCmd.Flags().StringS("t", "t", "", "specify a target pane")
	rootCmd.AddCommand(selectLayoutCmd)

	carapace.Gen(selectLayoutCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionPanes(),
	})

	carapace.Gen(selectLayoutCmd).PositionalCompletion(
		tmux.ActionLayouts(),
	)
}
