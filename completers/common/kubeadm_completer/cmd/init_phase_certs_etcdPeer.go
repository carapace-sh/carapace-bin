package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var init_phase_certs_etcdPeerCmd = &cobra.Command{
	Use:   "etcd-peer",
	Short: "Generate the certificate for etcd nodes to communicate with each other",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_certs_etcdPeerCmd).Standalone()

	init_phase_certs_etcdPeerCmd.Flags().String("cert-dir", "", "The path where to save and store the certificates.")
	init_phase_certs_etcdPeerCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_certs_etcdPeerCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_certs_etcdPeerCmd.Flags().String("kubernetes-version", "", "Choose a specific Kubernetes version for the control plane.")
	init_phase_certsCmd.AddCommand(init_phase_certs_etcdPeerCmd)

	carapace.Gen(init_phase_certs_etcdPeerCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
	})
}
