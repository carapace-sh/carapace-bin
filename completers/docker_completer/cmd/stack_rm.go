package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stack_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_rmCmd).Standalone()

	stack_rmCmd.Flags().String("orchestrator", "", "Orchestrator to use (swarm|kubernetes|all)")
	stackCmd.AddCommand(stack_rmCmd)

	carapace.Gen(stack_rmCmd).FlagCompletion(carapace.ActionMap{
		"orchestrator": carapace.ActionValues("swarm", "kubernetes", "all"),
	})

	carapace.Gen(stack_rmCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(args []string) carapace.Action {
			if stack_rmCmd.Flag("orchestrator").Changed {
				return action.ActionStacks(stack_rmCmd.Flag("orchestrator").Value.String()).Callback(args)
			} else {
				return action.ActionStacks("all").Callback(args)
			}
		}),
	)
}
