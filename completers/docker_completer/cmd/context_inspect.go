package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var context_inspectCmd = &cobra.Command{
	Use:   "inspect [OPTIONS] [CONTEXT] [CONTEXT...]",
	Short: "Display detailed information on one or more contexts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_inspectCmd).Standalone()

	context_inspectCmd.Flags().StringP("format", "f", "", "Format output using a custom template:")
	contextCmd.AddCommand(context_inspectCmd)

	carapace.Gen(context_inspectCmd).PositionalAnyCompletion(docker.ActionContexts())
}
