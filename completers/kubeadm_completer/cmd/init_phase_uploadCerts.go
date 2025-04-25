package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var init_phase_uploadCertsCmd = &cobra.Command{
	Use:   "upload-certs",
	Short: "Upload certificates to kubeadm-certs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(init_phase_uploadCertsCmd).Standalone()

	init_phase_uploadCertsCmd.Flags().String("certificate-key", "", "Key used to encrypt the control-plane certificates in the kubeadm-certs Secret. The certificate key is a hex encoded string that is an AES key of size 32 bytes.")
	init_phase_uploadCertsCmd.Flags().String("config", "", "Path to a kubeadm configuration file.")
	init_phase_uploadCertsCmd.Flags().Bool("dry-run", false, "Don't apply any changes; just output what would be done.")
	init_phase_uploadCertsCmd.Flags().String("kubeconfig", "", "The kubeconfig file to use when talking to the cluster. If the flag is not set, a set of standard locations can be searched for an existing kubeconfig file.")
	init_phase_uploadCertsCmd.Flags().Bool("skip-certificate-key-print", false, "Don't print the key used to encrypt the control-plane certificates.")
	init_phase_uploadCertsCmd.Flags().Bool("upload-certs", false, "Upload control-plane certificates to the kubeadm-certs Secret.")
	init_phaseCmd.AddCommand(init_phase_uploadCertsCmd)

	carapace.Gen(init_phase_uploadCertsCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(),
		"kubeconfig": carapace.ActionFiles(),
	})
}
