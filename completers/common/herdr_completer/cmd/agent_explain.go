package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var agent_explainCmd = &cobra.Command{
	Use:   "explain [target]",
	Short: "Explain agent detection state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(agent_explainCmd).Standalone()

	agent_explainCmd.Flags().String("agent", "", "")
	agent_explainCmd.Flags().String("file", "", "")
	agent_explainCmd.Flags().String("format", "", "")
	agent_explainCmd.Flags().Bool("json", false, "")
	agent_explainCmd.Flags().BoolP("verbose", "v", false, "")
	agentCmd.AddCommand(agent_explainCmd)

	carapace.Gen(agent_explainCmd).PositionalCompletion(herdr.ActionAgents())

	carapace.Gen(agent_explainCmd).FlagCompletion(carapace.ActionMap{
		"file":   carapace.ActionFiles(),
		"format": carapace.ActionValues("text", "json"),
	})
}
