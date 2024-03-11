package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kube_credentialsCmd = &cobra.Command{
	Use:    "credentials",
	Short:  "Get credentials for kubectl access",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kube_credentialsCmd).Standalone()

	kube_credentialsCmd.Flags().String("kube-cluster", "", "Name of the Kubernetes cluster to get credentials for.")
	kube_credentialsCmd.Flags().String("teleport-cluster", "", "Name of the Teleport cluster to get credentials for.")
	kube_credentialsCmd.MarkFlagRequired("kube-cluster")
	kube_credentialsCmd.MarkFlagRequired("teleport-cluster")
	kubeCmd.AddCommand(kube_credentialsCmd)
}
