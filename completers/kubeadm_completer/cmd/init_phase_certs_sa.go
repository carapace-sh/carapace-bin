package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var init_phase_certs_saCmd = &cobra.Command{
	Use:   "sa",
	Short: "Generate a private key for signing service account tokens along with its public key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_certs_saCmd).Standalone()

	init_phase_certs_saCmd.Flags().String("cert-dir", "", "The path where to save and store the certificates.")
	init_phase_certs_saCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_certs_saCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_certs_saCmd.Flags().String("kubernetes-version", "", "Choose a specific Kubernetes version for the control plane.")
	init_phase_certsCmd.AddCommand(init_phase_certs_saCmd)

	carapace.Gen(init_phase_certs_saCmd).FlagCompletion(carapace.ActionMap{
		"cert-dir": carapace.ActionDirectories(),
		"config":   carapace.ActionFiles(),
	})
}
