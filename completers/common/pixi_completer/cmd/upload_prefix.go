package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_prefixCmd = &cobra.Command{
	Use:   "prefix",
	Short: "Upload to a prefix.dev server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_prefixCmd).Standalone()

	upload_prefixCmd.Flags().StringP("api-key", "a", "", "The prefix.dev API key")
	upload_prefixCmd.Flags().String("attestation", "", "The path to the attestation file")
	upload_prefixCmd.Flags().StringP("channel", "c", "", "The channel to upload the package to")
	upload_prefixCmd.Flags().Bool("force", false, "Force upload even if the package already exists")
	upload_prefixCmd.Flags().Bool("generate-attestation", false, "Generate an attestation for the package")
	upload_prefixCmd.Flags().BoolP("skip-existing", "s", false, "Skip uploading if the package already exists")
	upload_prefixCmd.Flags().Bool("store-github-attestation", false, "Store the attestation on GitHub")
	upload_prefixCmd.Flags().StringP("url", "u", "", "The URL to the prefix.dev server")
	uploadCmd.AddCommand(upload_prefixCmd)

	carapace.Gen(upload_prefixCmd).FlagCompletion(carapace.ActionMap{
		"attestation": carapace.ActionFiles(),
	})
}
