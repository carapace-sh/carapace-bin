package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pane_resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "Resize a pane split",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_resizeCmd).Standalone()

	pane_resizeCmd.Flags().String("amount", "", "")
	pane_resizeCmd.Flags().Bool("current", false, "")
	pane_resizeCmd.Flags().String("direction", "", "")
	pane_resizeCmd.Flags().String("pane", "", "")
	paneCmd.AddCommand(pane_resizeCmd)

	carapace.Gen(pane_resizeCmd).FlagCompletion(carapace.ActionMap{
		"direction": carapace.ActionValues("left", "right", "up", "down"),
	})
}
