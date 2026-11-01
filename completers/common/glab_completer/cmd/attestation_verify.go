package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var attestation_verifyCmd = &cobra.Command{
	Use:   "verify <project-id> <artifact-path>",
	Short: "Verify the provenance of a specific artifact or file. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attestation_verifyCmd).Standalone()

	attestationCmd.AddCommand(attestation_verifyCmd)
}
