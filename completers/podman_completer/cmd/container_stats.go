package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_statsCmd = &cobra.Command{
	Use:   "stats [options] [CONTAINER...]",
	Short: "Display a live stream of container resource usage statistics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_statsCmd).Standalone()

	container_statsCmd.Flags().BoolP("all", "a", false, "Show all containers. Only running containers are shown by default. The default is false")
	container_statsCmd.Flags().String("format", "", "Pretty-print container statistics to JSON or using a Go template")
	container_statsCmd.Flags().StringP("interval", "i", "", "Time in seconds between stats reports")
	container_statsCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	container_statsCmd.Flags().Bool("no-reset", false, "Disable resetting the screen between intervals")
	container_statsCmd.Flags().Bool("no-stream", false, "Disable streaming stats and only pull the first result, default setting is false")
	container_statsCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	containerCmd.AddCommand(container_statsCmd)
}
