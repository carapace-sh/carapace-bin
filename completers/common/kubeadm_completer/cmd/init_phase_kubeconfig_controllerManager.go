package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var init_phase_kubeconfig_controllerManagerCmd = &cobra.Command{
	Use:   "controller-manager",
	Short: "Generate a kubeconfig file for the controller manager to use",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_kubeconfig_controllerManagerCmd).Standalone()

	init_phase_kubeconfig_controllerManagerCmd.Flags().String("apiserver-advertise-address", "", "The IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	init_phase_kubeconfig_controllerManagerCmd.Flags().String("apiserver-bind-port", "", "Port for the API Server to bind to.")
	init_phase_kubeconfig_controllerManagerCmd.Flags().String("cert-dir", "", "The path where to save and store the certificates.")
	init_phase_kubeconfig_controllerManagerCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_kubeconfig_controllerManagerCmd.Flags().String("control-plane-endpoint", "", "Specify a stable IP address or DNS name for the control plane.")
	init_phase_kubeconfig_controllerManagerCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_kubeconfig_controllerManagerCmd.Flags().String("kubeconfig-dir", "", "The path where to save the kubeconfig file.")
	init_phase_kubeconfig_controllerManagerCmd.Flags().String("kubernetes-version", "", "Choose a specific Kubernetes version for the control plane.")
	init_phase_kubeconfigCmd.AddCommand(init_phase_kubeconfig_controllerManagerCmd)

	carapace.Gen(init_phase_kubeconfig_controllerManagerCmd).FlagCompletion(carapace.ActionMap{
		"apiserver-bind-port": net.ActionPorts(),
		"cert-dir":            carapace.ActionDirectories(),
		"config":              carapace.ActionFiles(),
		"kubeconfig-dir":      carapace.ActionDirectories(),
	})
}
