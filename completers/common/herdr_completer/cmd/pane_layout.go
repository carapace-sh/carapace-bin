package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var pane_layoutCmd = &cobra.Command{
	Use:   "layout",
	Short: "Show pane layout information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_layoutCmd).Standalone()

	pane_layoutCmd.Flags().Bool("current", false, "")
	pane_layoutCmd.Flags().String("pane", "", "")
	paneCmd.AddCommand(pane_layoutCmd)

	carapace.Gen(pane_layoutCmd).FlagCompletion(carapace.ActionMap{
		"pane": herdr.ActionPanes(herdr.PaneOpts{}),
	})
}
