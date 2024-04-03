package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var attestation_inspectCmd = &cobra.Command{
	Use:    "inspect [<file path> | oci://<OCI image URI>] --bundle <path-to-bundle>",
	Short:  "Inspect a sigstore bundle",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attestation_inspectCmd).Standalone()

	attestation_inspectCmd.Flags().StringP("bundle", "b", "", "Path to bundle on disk, either a single bundle in a JSON file or a JSON lines file with multiple bundles")
	attestation_inspectCmd.Flags().StringP("digest-alg", "d", "", "The algorithm used to compute a digest of the artifact: {sha256|sha512}")
	attestation_inspectCmd.Flags().String("format", "", "Output format: {json}")
	attestation_inspectCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	attestation_inspectCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	attestation_inspectCmd.MarkFlagRequired("bundle")
	attestationCmd.AddCommand(attestation_inspectCmd)

	carapace.Gen(attestation_inspectCmd).FlagCompletion(carapace.ActionMap{
		"bundle":     carapace.ActionFiles(),
		"digest-alg": carapace.ActionValues("sha256", "sha512"),
		"format":     carapace.ActionValues("json"),
	})

	carapace.Gen(attestation_inspectCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
