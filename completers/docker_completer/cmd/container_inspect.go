package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/docker"
	"github.com/spf13/cobra"
)

var container_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Display detailed information on one or more containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_inspectCmd).Standalone()

	container_inspectCmd.Flags().StringP("format", "f", "", "Format the output using the given Go template")
	container_inspectCmd.Flags().BoolP("size", "s", false, "Display total file sizes")
	containerCmd.AddCommand(container_inspectCmd)

	carapace.Gen(container_inspectCmd).PositionalCompletion(
		docker.ActionContainers(),
	)
}
