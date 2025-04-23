package cmd

import (
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
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
	attestation_verifyCmd.Flags().Bool("bundle-from-oci", false, "When verifying an OCI image, fetch the attestation bundle from the OCI registry instead of from GitHub")
	attestation_verifyCmd.Flags().String("cert-identity", "", "Enforce that the certificate's SubjectAlternativeName matches the provided value exactly")
	attestation_verifyCmd.Flags().StringP("cert-identity-regex", "i", "", "Enforce that the certificate's SubjectAlternativeName matches the provided regex")
	attestation_verifyCmd.Flags().String("cert-oidc-issuer", "", "Enforce that the issuer of the OIDC token matches the provided value")
	attestation_verifyCmd.Flags().String("custom-trusted-root", "", "Path to a trusted_root.jsonl file; likely for offline verification")
	attestation_verifyCmd.Flags().Bool("deny-self-hosted-runners", false, "Fail verification for attestations generated on self-hosted runners")
	attestation_verifyCmd.Flags().StringP("digest-alg", "d", "", "The algorithm used to compute a digest of the artifact: {sha256|sha512}")
	attestation_verifyCmd.Flags().String("format", "", "Output format: {json}")
	attestation_verifyCmd.Flags().String("hostname", "", "Configure host to use")
	attestation_verifyCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	attestation_verifyCmd.Flags().StringP("limit", "L", "", "Maximum number of attestations to fetch")
	attestation_verifyCmd.Flags().Bool("no-public-good", false, "Do not verify attestations signed with Sigstore public good instance")
	attestation_verifyCmd.Flags().StringP("owner", "o", "", "GitHub organization to scope attestation lookup by")
	attestation_verifyCmd.Flags().String("predicate-type", "", "Enforce that verified attestations' predicate type matches the provided value")
	attestation_verifyCmd.Flags().StringP("repo", "R", "", "Repository name in the format <owner>/<repo>")
	attestation_verifyCmd.Flags().String("signer-digest", "", "Enforce that the digest associated with the signer workflow matches the provided value")
	attestation_verifyCmd.Flags().String("signer-repo", "", "Enforce that the workflow that signed the attestation's repository matches the provided value (<owner>/<repo>)")
	attestation_verifyCmd.Flags().String("signer-workflow", "", "Enforce that the workflow that signed the attestation matches the provided value ([host/]<owner>/<repo>/<path>/<to>/<workflow>)")
	attestation_verifyCmd.Flags().String("source-digest", "", "Enforce that the digest associated with the source repository matches the provided value")
	attestation_verifyCmd.Flags().String("source-ref", "", "Enforce that the git ref associated with the source repository matches the provided value")
	attestation_verifyCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	attestationCmd.AddCommand(attestation_verifyCmd)

	carapace.Gen(attestation_verifyCmd).FlagCompletion(carapace.ActionMap{
		"bundle":              carapace.ActionFiles(),
		"custom-trusted-root": carapace.ActionFiles(),
		"digest-alg":          carapace.ActionValues("sha256", "sha512"),
		"format":              carapace.ActionValues("json"),
		"hostname":            gh.ActionConfigHosts(),
		"owner":               gh.ActionOrganizations(gh.HostOpts{}),
		"repo":                gh.ActionHostOwnerRepositories(),
		"signer-repo":         gh.ActionHostOwnerRepositories(),
		"signer-workflow": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			n := 3
			if strings.Contains(strings.Split(c.Value, "/")[0], ".") {
				n = 4
			}

			if count := strings.Count(c.Value, "/"); count < n-1 {
				if count == n-2 {
					return gh.ActionHostOwnerRepositories().NoSpace().Suffix("/")
				}
				return gh.ActionHostOwnerRepositories().NoSpace()
			}

			return carapace.ActionMultiPartsN("/", n, func(c carapace.Context) carapace.Action {
				opts := gh.ContentOpts{
					Owner: c.Parts[n-3],
					Name:  c.Parts[n-2],
				}
				if n == 4 {
					opts.Host = c.Parts[0]
				}
				opts.Path = filepath.Dir(c.Value)
				prefix := filepath.Dir(c.Value)
				switch {
				case prefix == ".":
					prefix = ""
				default:
					prefix += "/"
				}
				return gh.ActionContents(opts).Prefix(prefix)
			})
		}),
		"source-ref": action.ActionBranches(attestation_verifyCmd), // TODO verify
	})

	carapace.Gen(attestation_verifyCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
