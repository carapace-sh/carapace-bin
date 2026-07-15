package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var agent_renameCmd = &cobra.Command{
	Use:   "rename <target> [name]",
	Short: "Rename an agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_renameCmd).Standalone()

	agent_renameCmd.Flags().Bool("clear", false, "")
	agentCmd.AddCommand(agent_renameCmd)

	carapace.Gen(agent_renameCmd).PositionalCompletion(action.ActionAgents())
}
