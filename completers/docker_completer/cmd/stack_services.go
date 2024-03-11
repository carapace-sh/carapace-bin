package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var stack_servicesCmd = &cobra.Command{
	Use:   "services [OPTIONS] STACK",
	Short: "List the services in the stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_servicesCmd).Standalone()

	stack_servicesCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	stack_servicesCmd.Flags().String("format", "", "Format output using a custom template:")
	stack_servicesCmd.Flags().BoolP("quiet", "q", false, "Only display IDs")
	stackCmd.AddCommand(stack_servicesCmd)

	carapace.Gen(stack_servicesCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if stack_servicesCmd.Flag("orchestrator").Changed {
				return docker.ActionStacks(stack_servicesCmd.Flag("orchestrator").Value.String())
			} else {
				return docker.ActionStacks("all")
			}
		}),
	)
}
