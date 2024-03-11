package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var context_updateCmd = &cobra.Command{
	Use:   "update [OPTIONS] CONTEXT",
	Short: "Update a context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_updateCmd).Standalone()

	context_updateCmd.Flags().String("description", "", "Description of the context")
	context_updateCmd.Flags().String("docker", "", "set the docker endpoint")
	contextCmd.AddCommand(context_updateCmd)

	carapace.Gen(context_updateCmd).PositionalCompletion(
		docker.ActionContexts(),
	)
}
