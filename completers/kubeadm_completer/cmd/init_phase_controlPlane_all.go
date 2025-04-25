package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubeadm"
	"github.com/spf13/cobra"
)

var init_phase_controlPlane_allCmd = &cobra.Command{
	Use:   "all",
	Short: "Generate all static Pod manifest files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_controlPlane_allCmd).Standalone()

	init_phase_controlPlane_allCmd.Flags().String("apiserver-advertise-address", "", "The IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	init_phase_controlPlane_allCmd.Flags().String("apiserver-bind-port", "", "Port for the API Server to bind to.")
	init_phase_controlPlane_allCmd.Flags().String("apiserver-extra-args", "", "A set of extra flags to pass to the API Server or override default ones in form of <flagname>=<value>")
	init_phase_controlPlane_allCmd.Flags().String("cert-dir", "", "The path where to save and store the certificates.")
	init_phase_controlPlane_allCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_controlPlane_allCmd.Flags().String("control-plane-endpoint", "", "Specify a stable IP address or DNS name for the control plane.")
	init_phase_controlPlane_allCmd.Flags().String("controller-manager-extra-args", "", "A set of extra flags to pass to the Controller Manager or override default ones in form of <flagname>=<value>")
	init_phase_controlPlane_allCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_controlPlane_allCmd.Flags().String("feature-gates", "", "A set of key=value pairs that describe feature gates for various features. Options are:")
	init_phase_controlPlane_allCmd.Flags().String("image-repository", "", "Choose a container registry to pull control plane images from")
	init_phase_controlPlane_allCmd.Flags().String("kubernetes-version", "", "Choose a specific Kubernetes version for the control plane.")
	init_phase_controlPlane_allCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\". For example, \"kube-apiserver0+merge.yaml\" or just \"etcd.json\". \"target\" can be one of \"kube-apiserver\", \"kube-controller-manager\", \"kube-scheduler\", \"etcd\", \"kubeletconfiguration\", \"corednsdeployment\". \"patchtype\" can be one of \"strategic\", \"merge\" or \"json\" and they match the patch formats supported by kubectl. The default \"patchtype\" is \"strategic\". \"extension\" must be either \"json\" or \"yaml\". \"suffix\" is an optional string that can be used to determine which patches are applied first alpha-numerically.")
	init_phase_controlPlane_allCmd.Flags().String("pod-network-cidr", "", "Specify range of IP addresses for the pod network. If set, the control plane will automatically allocate CIDRs for every node.")
	init_phase_controlPlane_allCmd.Flags().String("scheduler-extra-args", "", "A set of extra flags to pass to the Scheduler or override default ones in form of <flagname>=<value>")
	init_phase_controlPlane_allCmd.Flags().String("service-cidr", "", "Use alternative range of IP address for service VIPs.")
	init_phase_controlPlane_allCmd.Flag("apiserver-extra-args").Hidden = true
	init_phase_controlPlane_allCmd.Flag("controller-manager-extra-args").Hidden = true
	init_phase_controlPlane_allCmd.Flag("scheduler-extra-args").Hidden = true
	init_phase_controlPlaneCmd.AddCommand(init_phase_controlPlane_allCmd)

	carapace.Gen(init_phase_controlPlane_allCmd).FlagCompletion(carapace.ActionMap{
		"apiserver-bind-port": net.ActionPorts(),
		"cert-dir":            carapace.ActionDirectories(),
		"config":              carapace.ActionFiles(),
		"feature-gates":       kubeadm.ActionFeatureGates(),
		"patches":             carapace.ActionDirectories(),
	})
}
