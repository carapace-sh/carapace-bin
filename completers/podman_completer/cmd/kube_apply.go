package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kube_applyCmd = &cobra.Command{
	Use:   "apply [options] [CONTAINER...|POD...|VOLUME...]",
	Short: "Deploy a podman container, pod, volume, or Kubernetes yaml to a Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kube_applyCmd).Standalone()

	kube_applyCmd.Flags().String("ca-cert-file", "", "Path to the CA cert file for the Kubernetes cluster.")
	kube_applyCmd.Flags().StringP("file", "f", "", "Path to the Kubernetes yaml file to deploy.")
	kube_applyCmd.Flags().StringP("kubeconfig", "k", "", "Path to the kubeconfig file for the Kubernetes cluster")
	kube_applyCmd.Flags().String("ns", "", "The namespace to deploy the workload to on the Kubernetes cluster")
	kube_applyCmd.Flags().BoolP("service", "s", false, "Create a service object for the container being deployed.")
	kubeCmd.AddCommand(kube_applyCmd)
}
