package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var chooseTreeCmd = &cobra.Command{
	Use:   "choose-tree",
	Short: "put a window into tree choice mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chooseTreeCmd).Standalone()

	chooseTreeCmd.Flags().StringS("F", "F", "", "specify format for each list item")
	chooseTreeCmd.Flags().BoolS("G", "G", false, "include all sessions in any session groups")
	chooseTreeCmd.Flags().StringS("K", "K", "", "specify format for each shortcut key")
	chooseTreeCmd.Flags().BoolS("N", "N", false, "start without the preview")
	chooseTreeCmd.Flags().StringS("O", "O", "", "initial sort order")
	chooseTreeCmd.Flags().BoolS("Z", "Z", false, "zoom the pane")
	chooseTreeCmd.Flags().StringS("f", "f", "", "filter items")
	chooseTreeCmd.Flags().BoolS("h", "h", false, "hide the pane containing the mode")
	chooseTreeCmd.Flags().BoolS("i", "i", false, "start showing client information instead of the preview")
	chooseTreeCmd.Flags().BoolS("k", "k", false, "kill the pane when the mode is exited")
	chooseTreeCmd.Flags().BoolS("r", "r", false, "reverse sort order")
	chooseTreeCmd.Flags().BoolS("s", "s", false, "start with sessions collapsed")
	chooseTreeCmd.Flags().StringS("t", "t", "", "specify target pane")
	chooseTreeCmd.Flags().BoolS("w", "w", false, "start with windows collapsed")
	chooseTreeCmd.Flags().BoolS("y", "y", false, "disable confirmation prompts")
	rootCmd.AddCommand(chooseTreeCmd)

	carapace.Gen(chooseTreeCmd).FlagCompletion(carapace.ActionMap{
		"O": carapace.ActionValues("index", "name", "activity", "z"),
		"t": tmux.ActionPanes(),
	})
}
