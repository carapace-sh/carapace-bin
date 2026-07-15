package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pane_zoomCmd = &cobra.Command{
	Use:   "zoom",
	Short: "Toggle or set pane zoom",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_zoomCmd).Standalone()

	pane_zoomCmd.Flags().Bool("current", false, "")
	pane_zoomCmd.Flags().Bool("off", false, "")
	pane_zoomCmd.Flags().Bool("on", false, "")
	pane_zoomCmd.Flags().String("pane", "", "")
	pane_zoomCmd.Flags().Bool("toggle", false, "")
	paneCmd.AddCommand(pane_zoomCmd)
}
