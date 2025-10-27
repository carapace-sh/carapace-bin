package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stackCmd = &cobra.Command{
	Use:     "stack [OPTIONS]",
	Short:   "Manage Swarm stacks",
	GroupID: "swarm",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stackCmd).Standalone()

	stackCmd.PersistentFlags().String("orchestrator", "", "Orchestrator to use (swarm|all)")
	stackCmd.Flag("orchestrator").Hidden = true
	rootCmd.AddCommand(stackCmd)

	carapace.Gen(stackCmd).FlagCompletion(carapace.ActionMap{
		"orchestrator": carapace.ActionValues("swarm", "all"),
	})
}
