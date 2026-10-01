package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var copyModeCmd = &cobra.Command{
	Use:   "copy-mode",
	Short: "enter copy mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(copyModeCmd).Standalone()

	copyModeCmd.Flags().BoolS("H", "H", false, "hide the position indicator")
	copyModeCmd.Flags().BoolS("M", "M", false, "begin a mouse drag")
	copyModeCmd.Flags().BoolS("S", "S", false, "enter copy mode and scroll when bound to a mouse drag event")
	copyModeCmd.Flags().BoolS("d", "d", false, "enter copy mode and scroll one page down")
	copyModeCmd.Flags().BoolS("e", "e", false, "exit copy mode when scrolling to the bottom")
	copyModeCmd.Flags().BoolS("k", "k", false, "kill the pane when the mode is exited")
	copyModeCmd.Flags().BoolS("q", "q", false, "cancel copy mode and any other modes")
	copyModeCmd.Flags().StringS("s", "s", "", "specify source pane")
	copyModeCmd.Flags().StringS("t", "t", "", "specify target pane")
	copyModeCmd.Flags().BoolS("u", "u", false, "enter copy mode and scroll one page up")
	rootCmd.AddCommand(copyModeCmd)

	carapace.Gen(copyModeCmd).FlagCompletion(carapace.ActionMap{
		"s": tmux.ActionPanes(),
		"t": tmux.ActionPanes(),
	})
}
