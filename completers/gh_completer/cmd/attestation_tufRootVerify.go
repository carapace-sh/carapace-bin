package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var attestation_tufRootVerifyCmd = &cobra.Command{
	Use:    "tuf-root-verify --mirror <mirror-url> --root <root.json>",
	Short:  "Verify the TUF repository from a provided TUF root",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attestation_tufRootVerifyCmd).Standalone()

	attestation_tufRootVerifyCmd.Flags().StringP("mirror", "m", "", "URL to the TUF repository mirror")
	attestation_tufRootVerifyCmd.Flags().StringP("root", "r", "", "Path to the TUF root file on disk")
	attestation_tufRootVerifyCmd.MarkFlagRequired("mirror")
	attestation_tufRootVerifyCmd.MarkFlagRequired("root")
	attestationCmd.AddCommand(attestation_tufRootVerifyCmd)

	carapace.Gen(attestation_tufRootVerifyCmd).FlagCompletion(carapace.ActionMap{
		"root": carapace.ActionFiles(),
	})
}
