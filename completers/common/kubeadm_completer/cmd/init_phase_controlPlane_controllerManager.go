package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var init_phase_controlPlane_controllerManagerCmd = &cobra.Command{
	Use:   "controller-manager",
	Short: "Generates the kube-controller-manager static Pod manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_controlPlane_controllerManagerCmd).Standalone()

	init_phase_controlPlane_controllerManagerCmd.Flags().String("cert-dir", "", "The path where to save and store the certificates.")
	init_phase_controlPlane_controllerManagerCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_controlPlane_controllerManagerCmd.Flags().String("controller-manager-extra-args", "", "A set of extra flags to pass to the Controller Manager or override default ones in form of <flagname>=<value>")
	init_phase_controlPlane_controllerManagerCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_controlPlane_controllerManagerCmd.Flags().String("image-repository", "", "Choose a container registry to pull control plane images from")
	init_phase_controlPlane_controllerManagerCmd.Flags().String("kubernetes-version", "", "Choose a specific Kubernetes version for the control plane.")
	init_phase_controlPlane_controllerManagerCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\". For example, \"kube-apiserver0+merge.yaml\" or just \"etcd.json\". \"target\" can be one of \"kube-apiserver\", \"kube-controller-manager\", \"kube-scheduler\", \"etcd\", \"kubeletconfiguration\", \"corednsdeployment\". \"patchtype\" can be one of \"strategic\", \"merge\" or \"json\" and they match the patch formats supported by kubectl. The default \"patchtype\" is \"strategic\". \"extension\" must be either \"json\" or \"yaml\". \"suffix\" is an optional string that can be used to determine which patches are applied first alpha-numerically.")
	init_phase_controlPlane_controllerManagerCmd.Flags().String("pod-network-cidr", "", "Specify range of IP addresses for the pod network. If set, the control plane will automatically allocate CIDRs for every node.")
	init_phase_controlPlane_controllerManagerCmd.Flag("controller-manager-extra-args").Hidden = true
	init_phase_controlPlaneCmd.AddCommand(init_phase_controlPlane_controllerManagerCmd)

	carapace.Gen(init_phase_controlPlane_controllerManagerCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
		"patches":  carapace.ActionDirectories(),
	})
}
