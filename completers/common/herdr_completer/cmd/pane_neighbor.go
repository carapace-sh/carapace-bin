package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pane_neighborCmd = &cobra.Command{
	Use:   "neighbor",
	Short: "Find a pane neighbor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_neighborCmd).Standalone()

	pane_neighborCmd.Flags().Bool("current", false, "")
	pane_neighborCmd.Flags().String("direction", "", "")
	pane_neighborCmd.Flags().String("pane", "", "")
	paneCmd.AddCommand(pane_neighborCmd)

	carapace.Gen(pane_neighborCmd).FlagCompletion(carapace.ActionMap{
		"direction": carapace.ActionValues("left", "right", "up", "down"),
	})
}
