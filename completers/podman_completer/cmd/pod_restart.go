package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pod_restartCmd = &cobra.Command{
	Use:   "restart [options] POD [POD...]",
	Short: "Restart one or more pods",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_restartCmd).Standalone()

	pod_restartCmd.Flags().BoolP("all", "a", false, "Restart all running pods")
	pod_restartCmd.Flags().BoolP("latest", "l", false, "Act on the latest container podman is aware of")
	podCmd.AddCommand(pod_restartCmd)
}
