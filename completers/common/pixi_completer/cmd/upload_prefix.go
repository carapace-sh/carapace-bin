package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_prefixCmd = &cobra.Command{
	Use:   "prefix",
	Short: "Options for uploading to a prefix.dev server. Authentication is used from the keychain / auth-file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_prefixCmd).Standalone()

	upload_prefixCmd.Flags().StringP("api-key", "a", "", "The prefix.dev API key, if none is provided, the token is read from the keychain / auth-file")
	upload_prefixCmd.Flags().String("attestation", "", "Upload an attestation file alongside the package. Note: if you add an attestation, you can _only_ upload a single package. Mutually exclusive with --generate-attestation")
	upload_prefixCmd.Flags().StringP("channel", "c", "", "The channel to upload the package to")
	upload_prefixCmd.Flags().Bool("force", false, "Force overwrite existing packages")
	upload_prefixCmd.Flags().Bool("generate-attestation", false, "Automatically generate attestation using cosign in CI. Mutually exclusive with --attestation")
	upload_prefixCmd.Flags().BoolP("skip-existing", "s", false, "Skip upload if package already exists")
	upload_prefixCmd.Flags().Bool("store-github-attestation", false, "Also store the generated attestation to GitHub's attestation API. Requires `GITHUB_TOKEN` environment variable and only works in GitHub Actions. The attestation will be associated with the current repository")
	upload_prefixCmd.Flags().StringP("url", "u", "", "The URL to the prefix.dev server (only necessary for self-hosted instances)")
	upload_prefixCmd.MarkFlagRequired("channel")
	uploadCmd.AddCommand(upload_prefixCmd)

	carapace.Gen(upload_prefixCmd).FlagCompletion(carapace.ActionMap{
		"attestation": carapace.ActionFiles(),
	})
}
