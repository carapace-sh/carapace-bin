package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/herdr_completer/cmd/action"
	"github.com/spf13/cobra"
)

var agent_readCmd = &cobra.Command{
	Use:   "read <target>",
	Short: "Read agent terminal output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_readCmd).Standalone()

	agent_readCmd.Flags().Bool("ansi", false, "")
	agent_readCmd.Flags().String("format", "", "")
	agent_readCmd.Flags().String("lines", "", "")
	agent_readCmd.Flags().String("source", "", "")
	agentCmd.AddCommand(agent_readCmd)

	carapace.Gen(agent_readCmd).PositionalCompletion(action.ActionAgents())

	carapace.Gen(agent_readCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("text", "ansi"),
		"source": carapace.ActionValues("visible", "recent", "recent-unwrapped", "detection"),
	})
}
