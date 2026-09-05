package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var customizeModeCmd = &cobra.Command{
	Use:   "customize-mode",
	Short: "enter customize mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(customizeModeCmd).Standalone()

	customizeModeCmd.Flags().StringS("F", "F", "", "specify format for each item")
	customizeModeCmd.Flags().BoolS("N", "N", false, "start without the option information")
	customizeModeCmd.Flags().BoolS("Z", "Z", false, "zoom the pane")
	customizeModeCmd.Flags().StringS("f", "f", "", "specify initial filter")
	customizeModeCmd.Flags().BoolS("k", "k", false, "kill the pane when the mode is exited")
	customizeModeCmd.Flags().StringS("t", "t", "", "specify target pane")
	rootCmd.AddCommand(customizeModeCmd)

	carapace.Gen(customizeModeCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionPanes(),
	})
}
