package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_statsCmd = &cobra.Command{
	Use:   "stats [OPTIONS] [CONTAINER...]",
	Short: "Display a live stream of container(s) resource usage statistics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_statsCmd).Standalone()

	container_statsCmd.Flags().BoolP("all", "a", false, "Show all containers (default shows just running)")
	container_statsCmd.Flags().String("format", "", "Format output using a custom template:")
	container_statsCmd.Flags().Bool("no-stream", false, "Disable streaming stats and only pull the first result")
	container_statsCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	containerCmd.AddCommand(container_statsCmd)

	carapace.Gen(container_statsCmd).PositionalAnyCompletion(docker.ActionContainers())
}
