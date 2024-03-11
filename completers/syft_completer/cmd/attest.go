package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/syft"
	"github.com/spf13/cobra"
)

var attestCmd = &cobra.Command{
	Use:   "attest --output [FORMAT] <IMAGE>",
	Short: "Generate an SBOM as an attestation for the given [SOURCE] container image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attestCmd).Standalone()
	attestCmd.Flags().StringP("key", "k", "", "the key to use for the attestation")
	addCommonFlags(attestCmd)

	rootCmd.AddCommand(attestCmd)

	carapace.Gen(attestCmd).PositionalCompletion(
		syft.ActionSources(),
	)
}
