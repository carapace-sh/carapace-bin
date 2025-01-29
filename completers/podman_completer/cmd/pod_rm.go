package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pod_rmCmd = &cobra.Command{
	Use:   "rm [options] POD [POD...]",
	Short: "Remove one or more pods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_rmCmd).Standalone()

	pod_rmCmd.Flags().BoolP("all", "a", false, "Remove all running pods")
	pod_rmCmd.Flags().BoolP("force", "f", false, "Force removal of a running pod by first stopping all containers, then removing all containers in the pod.  The default is false")
	pod_rmCmd.Flags().BoolP("ignore", "i", false, "Ignore errors when a specified pod is missing")
	pod_rmCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	pod_rmCmd.Flags().StringSlice("pod-id-file", []string{}, "Read the pod ID from the file")
	pod_rmCmd.Flags().StringP("time", "t", "", "Seconds to wait for pod stop before killing the container")
	podCmd.AddCommand(pod_rmCmd)
}
