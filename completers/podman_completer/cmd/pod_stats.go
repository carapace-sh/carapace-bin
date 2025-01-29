package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pod_statsCmd = &cobra.Command{
	Use:   "stats [options] [POD...]",
	Short: "Display a live stream of resource usage statistics for the containers in one or more pods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_statsCmd).Standalone()

	pod_statsCmd.Flags().BoolP("all", "a", false, "Provide stats for all pods")
	pod_statsCmd.Flags().String("format", "", "Pretty-print container statistics to JSON or using a Go template")
	pod_statsCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	pod_statsCmd.Flags().Bool("no-reset", false, "Disable resetting the screen when streaming")
	pod_statsCmd.Flags().Bool("no-stream", false, "Disable streaming stats and only pull the first result")
	podCmd.AddCommand(pod_statsCmd)
}
