package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var pane_releaseAgentCmd = &cobra.Command{
	Use:   "release-agent <pane_id>",
	Short: "Release pane agent lifecycle authority",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_releaseAgentCmd).Standalone()

	pane_releaseAgentCmd.Flags().String("agent", "", "")
	pane_releaseAgentCmd.Flags().String("seq", "", "")
	pane_releaseAgentCmd.Flags().String("source", "", "")
	paneCmd.AddCommand(pane_releaseAgentCmd)

	carapace.Gen(pane_releaseAgentCmd).PositionalCompletion(herdr.ActionPanes(herdr.PaneOpts{}))
}
