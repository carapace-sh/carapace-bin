package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
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

	carapace.Gen(pane_releaseAgentCmd).PositionalCompletion(action.ActionPanes(pane_releaseAgentCmd))
}
