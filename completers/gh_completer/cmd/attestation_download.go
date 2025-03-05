package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var attestation_downloadCmd = &cobra.Command{
	Use:   "download [<file-path> | oci://<image-uri>] [--owner | --repo]",
	Short: "Download an artifact's attestations for offline use",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attestation_downloadCmd).Standalone()

	attestation_downloadCmd.Flags().StringP("digest-alg", "d", "", "The algorithm used to compute a digest of the artifact: {sha256|sha512}")
	attestation_downloadCmd.Flags().String("hostname", "", "Configure host to use")
	attestation_downloadCmd.Flags().StringP("limit", "L", "", "Maximum number of attestations to fetch")
	attestation_downloadCmd.Flags().StringP("owner", "o", "", "GitHub organization to scope attestation lookup by")
	attestation_downloadCmd.Flags().String("predicate-type", "", "Filter attestations by provided predicate type")
	attestation_downloadCmd.Flags().StringP("repo", "R", "", "Repository name in the format <owner>/<repo>")
	attestationCmd.AddCommand(attestation_downloadCmd)

	carapace.Gen(attestation_downloadCmd).FlagCompletion(carapace.ActionMap{
		"digest-alg": carapace.ActionValues("sha256", "sha512"),
		"hostname":   gh.ActionConfigHosts(),
		"owner":      gh.ActionOrganizations(gh.HostOpts{}),
		"repo":       gh.ActionHostOwnerRepositories(),
	})

	carapace.Gen(attestation_downloadCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
