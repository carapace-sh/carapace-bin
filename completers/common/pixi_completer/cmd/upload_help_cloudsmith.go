package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_help_cloudsmithCmd = &cobra.Command{
	Use:   "cloudsmith",
	Short: "Options for uploading to a Cloudsmith repository. Authentication is used from the keychain / auth-file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_help_cloudsmithCmd).Standalone()

	upload_helpCmd.AddCommand(upload_help_cloudsmithCmd)
}
