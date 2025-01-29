package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pod_stopCmd = &cobra.Command{
	Use:   "stop [options] POD [POD...]",
	Short: "Stop one or more pods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_stopCmd).Standalone()

	pod_stopCmd.Flags().BoolP("all", "a", false, "Stop all running pods")
	pod_stopCmd.Flags().BoolP("ignore", "i", false, "Ignore errors when a specified pod is missing")
	pod_stopCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	pod_stopCmd.Flags().StringSlice("pod-id-file", []string{}, "Write the pod ID to the file")
	pod_stopCmd.Flags().StringP("time", "t", "", "Seconds to wait for pod stop before killing the container")
	podCmd.AddCommand(pod_stopCmd)
}
