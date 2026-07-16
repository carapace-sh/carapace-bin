package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_cloudsmithCmd = &cobra.Command{
	Use:   "cloudsmith",
	Short: "Options for uploading to a Cloudsmith repository. Authentication is used from the keychain / auth-file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_cloudsmithCmd).Standalone()

	upload_cloudsmithCmd.Flags().StringP("api-key", "a", "", "The Cloudsmith API key, if none is provided, the token is read from the keychain / auth-file")
	upload_cloudsmithCmd.Flags().StringP("owner", "o", "", "The owner (namespace) of the Cloudsmith repository")
	upload_cloudsmithCmd.Flags().StringP("repo", "r", "", "The Cloudsmith repository name")
	upload_cloudsmithCmd.Flags().StringP("url", "u", "", "The URL to the Cloudsmith API server")
	upload_cloudsmithCmd.MarkFlagRequired("owner")
	upload_cloudsmithCmd.MarkFlagRequired("repo")
	uploadCmd.AddCommand(upload_cloudsmithCmd)
}
