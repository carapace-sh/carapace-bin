package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pod_pauseCmd = &cobra.Command{
	Use:   "pause [options] POD [POD...]",
	Short: "Pause one or more pods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_pauseCmd).Standalone()

	pod_pauseCmd.Flags().BoolP("all", "a", false, "Pause all running pods")
	pod_pauseCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	podCmd.AddCommand(pod_pauseCmd)
}
