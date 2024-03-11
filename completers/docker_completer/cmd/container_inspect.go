package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_inspectCmd = &cobra.Command{
	Use:   "inspect [OPTIONS] CONTAINER [CONTAINER...]",
	Short: "Display detailed information on one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_inspectCmd).Standalone()

	container_inspectCmd.Flags().StringP("format", "f", "", "Format output using a custom template:")
	container_inspectCmd.Flags().BoolP("size", "s", false, "Display total file sizes")
	containerCmd.AddCommand(container_inspectCmd)

	carapace.Gen(container_inspectCmd).PositionalAnyCompletion(
		docker.ActionContainers().FilterArgs(),
	)
}
