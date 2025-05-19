package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var init_phase_addon_kubeProxyCmd = &cobra.Command{
	Use:   "kube-proxy",
	Short: "Install the kube-proxy addon to a Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_addon_kubeProxyCmd).Standalone()

	init_phase_addon_kubeProxyCmd.Flags().String("apiserver-advertise-address", "", "The IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	init_phase_addon_kubeProxyCmd.Flags().String("apiserver-bind-port", "", "Port for the API Server to bind to.")
	init_phase_addon_kubeProxyCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_addon_kubeProxyCmd.Flags().String("control-plane-endpoint", "", "Specify a stable IP address or DNS name for the control plane.")
	init_phase_addon_kubeProxyCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_addon_kubeProxyCmd.Flags().String("image-repository", "", "Choose a container registry to pull control plane images from")
	init_phase_addon_kubeProxyCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	init_phase_addon_kubeProxyCmd.Flags().String("kubernetes-version", "", "Choose a specific Kubernetes version for the control plane.")
	init_phase_addon_kubeProxyCmd.Flags().String("pod-network-cidr", "", "Specify range of IP addresses for the pod network. If set, the control plane will automatically allocate CIDRs for every node.")
	init_phase_addon_kubeProxyCmd.Flags().Bool("print-manifest", false, "Print the addon manifests to STDOUT instead of installing them")
	init_phase_addonCmd.AddCommand(init_phase_addon_kubeProxyCmd)

	carapace.Gen(init_phase_addon_kubeProxyCmd).FlagCompletion(carapace.ActionMap{
		"apiserver-bind-port": net.ActionPorts(),
		"config":              carapace.ActionFiles(),
		"kubeconfig":          carapace.ActionFiles(),
	})
}
