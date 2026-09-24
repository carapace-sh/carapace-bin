package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var chooseClientCmd = &cobra.Command{
	Use:   "choose-client",
	Short: "put a window into client choice mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chooseClientCmd).Standalone()

	chooseClientCmd.Flags().StringS("F", "F", "", "specify format for each list item")
	chooseClientCmd.Flags().StringS("K", "K", "", "specify format for each shortcut key")
	chooseClientCmd.Flags().BoolS("N", "N", false, "start without the preview")
	chooseClientCmd.Flags().StringS("O", "O", "", "initial sort order")
	chooseClientCmd.Flags().BoolS("Z", "Z", false, "zoom the pane")
	chooseClientCmd.Flags().StringS("f", "f", "", "filter items")
	chooseClientCmd.Flags().BoolS("h", "h", false, "hide the pane containing the mode")
	chooseClientCmd.Flags().BoolS("k", "k", false, "kill the pane when the mode is exited")
	chooseClientCmd.Flags().BoolS("r", "r", false, "reverse sort order")
	chooseClientCmd.Flags().StringS("t", "t", "", "specify target pane")
	chooseClientCmd.Flags().BoolS("y", "y", false, "disable confirmation prompts")
	rootCmd.AddCommand(chooseClientCmd)

	carapace.Gen(chooseClientCmd).FlagCompletion(carapace.ActionMap{
		"O": carapace.ActionValues("name", "size", "creation", "activity"),
		"t": tmux.ActionPanes(),
	})
}
