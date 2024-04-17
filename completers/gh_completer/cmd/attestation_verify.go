package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var attestation_verifyCmd = &cobra.Command{
	Use:   "verify [<file-path> | oci://<image-uri>] [--owner | --repo]",
	Short: "Verify an artifact's integrity using attestations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attestation_verifyCmd).Standalone()

	attestation_verifyCmd.Flags().StringP("bundle", "b", "", "Path to bundle on disk, either a single bundle in a JSON file or a JSON lines file with multiple bundles")
	attestation_verifyCmd.Flags().String("cert-identity", "", "Enforce that the certificate's subject alternative name matches the provided value exactly")
	attestation_verifyCmd.Flags().StringP("cert-identity-regex", "i", "", "Enforce that the certificate's subject alternative name matches the provided regex")
	attestation_verifyCmd.Flags().String("cert-oidc-issuer", "", "Issuer of the OIDC token")
	attestation_verifyCmd.Flags().String("custom-trusted-root", "", "Path to a custom trustedroot.json file to use for verification")
	attestation_verifyCmd.Flags().Bool("deny-self-hosted-runners", false, "Fail verification for attestations generated on self-hosted runners.")
	attestation_verifyCmd.Flags().StringP("digest-alg", "d", "", "The algorithm used to compute a digest of the artifact: {sha256|sha512}")
	attestation_verifyCmd.Flags().String("format", "", "Output format: {json}")
	attestation_verifyCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	attestation_verifyCmd.Flags().StringP("limit", "L", "", "Maximum number of attestations to fetch")
	attestation_verifyCmd.Flags().Bool("no-public-good", false, "Only verify attestations signed with GitHub's Sigstore instance")
	attestation_verifyCmd.Flags().StringP("owner", "o", "", "GitHub organization to scope attestation lookup by")
	attestation_verifyCmd.Flags().String("predicate-type", "", "Filter attestations by provided predicate type")
	attestation_verifyCmd.Flags().StringP("repo", "R", "", "Repository name in the format <owner>/<repo>")
	attestation_verifyCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	attestationCmd.AddCommand(attestation_verifyCmd)

	carapace.Gen(attestation_verifyCmd).FlagCompletion(carapace.ActionMap{
		"bundle":              carapace.ActionFiles(),
		"custom-trusted-root": carapace.ActionFiles(),
		"digest-alg":          carapace.ActionValues("sha256", "sha512"),
		"format":              carapace.ActionValues("json"),
		"owner":               gh.ActionOrganizations(gh.HostOpts{}),
		"repo":                gh.ActionHostOwnerRepositories(),
	})

	carapace.Gen(attestation_verifyCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
