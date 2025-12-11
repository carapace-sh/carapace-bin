package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var init_phase_etcd_localCmd = &cobra.Command{
	Use:   "local",
	Short: "Generate the static Pod manifest file for a local, single-node local etcd instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_etcd_localCmd).Standalone()

	init_phase_etcd_localCmd.Flags().String("cert-dir", "", "The path where to save and store the certificates.")
	init_phase_etcd_localCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_etcd_localCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_etcd_localCmd.Flags().String("image-repository", "", "Choose a container registry to pull control plane images from")
	init_phase_etcd_localCmd.Flags().String("patches", "", "Path to a directory that contains files named \"target[suffix][+patchtype].extension\". For example, \"kube-apiserver0+merge.yaml\" or just \"etcd.json\". \"target\" can be one of \"kube-apiserver\", \"kube-controller-manager\", \"kube-scheduler\", \"etcd\", \"kubeletconfiguration\", \"corednsdeployment\". \"patchtype\" can be one of \"strategic\", \"merge\" or \"json\" and they match the patch formats supported by kubectl. The default \"patchtype\" is \"strategic\". \"extension\" must be either \"json\" or \"yaml\". \"suffix\" is an optional string that can be used to determine which patches are applied first alpha-numerically.")
	init_phase_etcdCmd.AddCommand(init_phase_etcd_localCmd)

	carapace.Gen(init_phase_etcd_localCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
		"patches":  carapace.ActionDirectories(),
	})
}
