package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_kubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "[ALPHA] Connect to a Kubernetes cluster using a Tailscale Auth Proxy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_kubeconfigCmd).Standalone()

	configureCmd.AddCommand(configure_kubeconfigCmd)
}
