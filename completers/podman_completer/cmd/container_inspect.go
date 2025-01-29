package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_inspectCmd = &cobra.Command{
	Use:   "inspect [options] CONTAINER [CONTAINER...]",
	Short: "Display the configuration of a container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_inspectCmd).Standalone()

	container_inspectCmd.Flags().StringP("format", "f", "", "Format the output to a Go template or json")
	container_inspectCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_inspectCmd.Flags().BoolP("size", "s", false, "Display total file size")
	containerCmd.AddCommand(container_inspectCmd)
}
