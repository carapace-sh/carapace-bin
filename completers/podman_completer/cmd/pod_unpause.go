package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pod_unpauseCmd = &cobra.Command{
	Use:   "unpause [options] POD [POD...]",
	Short: "Unpause one or more pods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_unpauseCmd).Standalone()

	pod_unpauseCmd.Flags().BoolP("all", "a", false, "Unpause all running pods")
	pod_unpauseCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	podCmd.AddCommand(pod_unpauseCmd)
}
