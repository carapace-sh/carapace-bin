package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var wait_agentStatusCmd = &cobra.Command{
	Use:   "agent-status <pane_id>",
	Short: "Wait for pane agent status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wait_agentStatusCmd).Standalone()

	wait_agentStatusCmd.Flags().String("status", "", "")
	wait_agentStatusCmd.Flags().String("timeout", "", "")
	wait_agentStatusCmd.MarkFlagRequired("status")
	waitCmd.AddCommand(wait_agentStatusCmd)

	carapace.Gen(wait_agentStatusCmd).PositionalCompletion(action.ActionPanes(wait_agentStatusCmd))

	carapace.Gen(wait_agentStatusCmd).FlagCompletion(carapace.ActionMap{
		"status": carapace.ActionValues("idle", "working", "blocked", "done", "unknown"),
	})
}
