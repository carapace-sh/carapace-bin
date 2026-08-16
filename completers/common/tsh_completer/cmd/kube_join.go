package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kube_joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Join an active Kubernetes session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kube_joinCmd).Standalone()

	kube_joinCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect.")
	kube_joinCmd.Flags().StringP("mode", "m", "observer", "Mode of joining the session.")
	kubeCmd.AddCommand(kube_joinCmd)
}
