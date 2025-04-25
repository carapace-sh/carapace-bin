package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubeadm"
	"github.com/spf13/cobra"
)

var init_phase_addon_corednsCmd = &cobra.Command{
	Use:   "coredns",
	Short: "Install the CoreDNS addon to a Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_addon_corednsCmd).Standalone()

	init_phase_addon_corednsCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_addon_corednsCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_addon_corednsCmd.Flags().String("feature-gates", "", "A set of key=value pairs that describe feature gates for various features. Options are:")
	init_phase_addon_corednsCmd.Flags().String("image-repository", "", "Choose a container registry to pull control plane images from")
	init_phase_addon_corednsCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	init_phase_addon_corednsCmd.Flags().String("kubernetes-version", "", "Choose a specific Kubernetes version for the control plane.")
	init_phase_addon_corednsCmd.Flags().Bool("print-manifest", false, "Print the addon manifests to STDOUT instead of installing them")
	init_phase_addon_corednsCmd.Flags().String("service-cidr", "", "Use alternative range of IP address for service VIPs.")
	init_phase_addon_corednsCmd.Flags().String("service-dns-domain", "", "Use alternative domain for services, e.g. \"myorg.internal\".")
	init_phase_addonCmd.AddCommand(init_phase_addon_corednsCmd)

	carapace.Gen(init_phase_addon_corednsCmd).FlagCompletion(carapace.ActionMap{
		"config":        carapace.ActionFiles(),
		"feature-gates": kubeadm.ActionFeatureGates(),
		"kubeconfig":    carapace.ActionFiles(),
	})
}
