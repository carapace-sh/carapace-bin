package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var compute_agent_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set compute agent properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_agent_setCmd).Standalone()

	compute_agent_setCmd.Flags().String("agent-version", "", "Version of the agent")
	compute_agent_setCmd.Flags().String("md5hash", "", "MD5 hash of the agent")
	compute_agent_setCmd.Flags().String("url", "", "URL of the agent")
	compute_agentCmd.AddCommand(compute_agent_setCmd)
}
