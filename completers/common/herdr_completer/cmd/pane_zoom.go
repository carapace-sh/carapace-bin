package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
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

	carapace.Gen(pane_zoomCmd).PositionalCompletion(herdr.ActionPanes(herdr.PaneOpts{}))

	carapace.Gen(pane_zoomCmd).FlagCompletion(carapace.ActionMap{
		"pane": herdr.ActionPanes(herdr.PaneOpts{}),
	})
}
