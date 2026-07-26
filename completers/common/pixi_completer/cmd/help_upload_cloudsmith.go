package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_upload_cloudsmithCmd = &cobra.Command{
	Use:   "cloudsmith",
	Short: "Options for uploading to a Cloudsmith repository. Authentication is used from the keychain / auth-file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_upload_cloudsmithCmd).Standalone()

	help_uploadCmd.AddCommand(help_upload_cloudsmithCmd)
}
