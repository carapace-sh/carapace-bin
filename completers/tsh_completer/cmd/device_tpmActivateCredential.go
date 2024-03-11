package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var device_tpmActivateCredentialCmd = &cobra.Command{
	Use:    "tpm-activate-credential",
	Short:  "",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_tpmActivateCredentialCmd).Standalone()

	device_tpmActivateCredentialCmd.Flags().String("encrypted-credential", "", "")
	device_tpmActivateCredentialCmd.Flags().String("encrypted-credential-secret", "", "")
	device_tpmActivateCredentialCmd.MarkFlagRequired("encrypted-credential")
	device_tpmActivateCredentialCmd.MarkFlagRequired("encrypted-credential-secret")
	deviceCmd.AddCommand(device_tpmActivateCredentialCmd)
}
