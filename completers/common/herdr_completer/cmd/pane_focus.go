package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pane_focusCmd = &cobra.Command{
	Use:   "focus",
	Short: "Focus a neighboring pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_focusCmd).Standalone()

	pane_focusCmd.Flags().Bool("current", false, "")
	pane_focusCmd.Flags().String("direction", "", "")
	pane_focusCmd.Flags().String("pane", "", "")
	paneCmd.AddCommand(pane_focusCmd)

	carapace.Gen(pane_focusCmd).FlagCompletion(carapace.ActionMap{
		"direction": carapace.ActionValues("left", "right", "up", "down"),
	})
}
