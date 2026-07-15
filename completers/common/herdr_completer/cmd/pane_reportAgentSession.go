package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pane_reportAgentSessionCmd = &cobra.Command{
	Use:   "report-agent-session <pane_id>",
	Short: "Report pane agent session identity",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_reportAgentSessionCmd).Standalone()

	pane_reportAgentSessionCmd.Flags().String("agent", "", "")
	pane_reportAgentSessionCmd.Flags().String("agent-session-id", "", "")
	pane_reportAgentSessionCmd.Flags().String("agent-session-path", "", "")
	pane_reportAgentSessionCmd.Flags().String("seq", "", "")
	pane_reportAgentSessionCmd.Flags().String("session-start-source", "", "")
	pane_reportAgentSessionCmd.Flags().String("source", "", "")
	paneCmd.AddCommand(pane_reportAgentSessionCmd)

	carapace.Gen(pane_reportAgentSessionCmd).PositionalCompletion(action.ActionPanes(pane_reportAgentSessionCmd))

	carapace.Gen(pane_reportAgentSessionCmd).FlagCompletion(carapace.ActionMap{
		"agent-session-path": carapace.ActionFiles(),
	})
}
