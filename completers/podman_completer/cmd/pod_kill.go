package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pod_killCmd = &cobra.Command{
	Use:   "kill [options] POD [POD...]",
	Short: "Send the specified signal or SIGKILL to containers in pod",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_killCmd).Standalone()

	pod_killCmd.Flags().BoolP("all", "a", false, "Kill all containers in all pods")
	pod_killCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	pod_killCmd.Flags().StringP("signal", "s", "", "Signal to send to the containers in the pod")
	podCmd.AddCommand(pod_killCmd)
}
