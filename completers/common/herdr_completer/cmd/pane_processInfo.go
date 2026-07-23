package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var pane_processInfoCmd = &cobra.Command{
	Use:   "process-info",
	Short: "Show pane process information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_processInfoCmd).Standalone()

	pane_processInfoCmd.Flags().Bool("current", false, "")
	pane_processInfoCmd.Flags().String("pane", "", "")
	paneCmd.AddCommand(pane_processInfoCmd)

	carapace.Gen(pane_processInfoCmd).FlagCompletion(carapace.ActionMap{
		"pane": herdr.ActionPanes(herdr.PaneOpts{}),
	})
}
