package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var selfHostedRunner_orchestratorCmd = &cobra.Command{
	Use:   "orchestrator",
	Short: "Spawn on-demand runners",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(selfHostedRunner_orchestratorCmd).Standalone()

	selfHostedRunnerCmd.AddCommand(selfHostedRunner_orchestratorCmd)
}
