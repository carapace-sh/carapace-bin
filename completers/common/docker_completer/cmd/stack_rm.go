package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var stack_rmCmd = &cobra.Command{
	Use:     "rm [OPTIONS] STACK [STACK...]",
	Short:   "Remove one or more stacks",
	Aliases: []string{"remove", "down"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_rmCmd).Standalone()

	stack_rmCmd.Flags().BoolP("detach", "d", false, "Do not wait for stack removal")
	stackCmd.AddCommand(stack_rmCmd)

	carapace.Gen(stack_rmCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if stack_rmCmd.Flag("orchestrator").Changed {
				return docker.ActionStacks(stack_rmCmd.Flag("orchestrator").Value.String())
			} else {
				return docker.ActionStacks("all")
			}
		}),
	)
}
