package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var stack_deployCmd = &cobra.Command{
	Use:     "deploy [OPTIONS] STACK",
	Short:   "Deploy a new stack or update an existing stack",
	Aliases: []string{"up"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_deployCmd).Standalone()

	stack_deployCmd.Flags().StringSliceP("compose-file", "c", nil, "Path to a Compose file, or \"-\" to read from stdin")
	stack_deployCmd.Flags().Bool("prune", false, "Prune services that are no longer referenced")
	stack_deployCmd.Flags().String("resolve-image", "always", "Query the registry to resolve image digest and supported platforms (\"always\", \"changed\", \"never\")")
	stack_deployCmd.Flags().Bool("with-registry-auth", false, "Send registry authentication details to Swarm agents")
	stackCmd.AddCommand(stack_deployCmd)

	carapace.Gen(stack_deployCmd).FlagCompletion(carapace.ActionMap{
		"compose-file":  carapace.ActionFiles(".yml"),
		"resolve-image": carapace.ActionValues("always", "changed", "never"),
	})

	carapace.Gen(stack_deployCmd).PositionalCompletion(
		docker.ActionStacks("all"),
	)
}
