package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pane_moveCmd = &cobra.Command{
	Use:   "move <pane_id>",
	Short: "Move a pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_moveCmd).Standalone()

	pane_moveCmd.Flags().Bool("focus", false, "")
	pane_moveCmd.Flags().String("label", "", "")
	pane_moveCmd.Flags().Bool("new-tab", false, "")
	pane_moveCmd.Flags().Bool("new-workspace", false, "")
	pane_moveCmd.Flags().Bool("no-focus", false, "")
	pane_moveCmd.Flags().String("ratio", "", "")
	pane_moveCmd.Flags().String("split", "", "")
	pane_moveCmd.Flags().String("tab", "", "")
	pane_moveCmd.Flags().String("tab-label", "", "")
	pane_moveCmd.Flags().String("target-pane", "", "")
	pane_moveCmd.Flags().String("workspace", "", "")
	paneCmd.AddCommand(pane_moveCmd)

	carapace.Gen(pane_moveCmd).FlagCompletion(carapace.ActionMap{
		"split": carapace.ActionValues("right", "down"),
	})
}
