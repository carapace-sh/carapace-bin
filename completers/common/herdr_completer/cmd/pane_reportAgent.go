package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pane_reportAgentCmd = &cobra.Command{
	Use:   "report-agent <pane_id>",
	Short: "Report pane agent lifecycle state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_reportAgentCmd).Standalone()

	pane_reportAgentCmd.Flags().String("agent", "", "")
	pane_reportAgentCmd.Flags().String("agent-session-id", "", "")
	pane_reportAgentCmd.Flags().String("agent-session-path", "", "")
	pane_reportAgentCmd.Flags().String("message", "", "")
	pane_reportAgentCmd.Flags().String("seq", "", "")
	pane_reportAgentCmd.Flags().String("source", "", "")
	pane_reportAgentCmd.Flags().String("state", "", "")
	pane_reportAgentCmd.MarkFlagRequired("state")
	paneCmd.AddCommand(pane_reportAgentCmd)

	carapace.Gen(pane_reportAgentCmd).PositionalCompletion(action.ActionPanes(pane_reportAgentCmd))

	carapace.Gen(pane_reportAgentCmd).FlagCompletion(carapace.ActionMap{
		"agent-session-path": carapace.ActionFiles(),
		"state":              carapace.ActionValues("idle", "working", "blocked", "unknown"),
	})
}
