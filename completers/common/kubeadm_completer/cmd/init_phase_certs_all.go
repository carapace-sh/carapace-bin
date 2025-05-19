package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var init_phase_certs_allCmd = &cobra.Command{
	Use:   "all",
	Short: "Generate all certificates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_certs_allCmd).Standalone()

	init_phase_certs_allCmd.Flags().String("apiserver-advertise-address", "", "The IP address the API Server will advertise it's listening on. If not set the default network interface will be used.")
	init_phase_certs_allCmd.Flags().StringSlice("apiserver-cert-extra-sans", nil, "Optional extra Subject Alternative Names (SANs) to use for the API Server serving certificate. Can be both IP addresses and DNS names.")
	init_phase_certs_allCmd.Flags().String("cert-dir", "", "The path where to save and store the certificates.")
	init_phase_certs_allCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_certs_allCmd.Flags().String("control-plane-endpoint", "", "Specify a stable IP address or DNS name for the control plane.")
	init_phase_certs_allCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_certs_allCmd.Flags().String("kubernetes-version", "", "Choose a specific Kubernetes version for the control plane.")
	init_phase_certs_allCmd.Flags().String("service-cidr", "", "Use alternative range of IP address for service VIPs.")
	init_phase_certs_allCmd.Flags().String("service-dns-domain", "", "Use alternative domain for services, e.g. \"myorg.internal\".")
	init_phase_certsCmd.AddCommand(init_phase_certs_allCmd)

	carapace.Gen(init_phase_certs_allCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
	})
}
