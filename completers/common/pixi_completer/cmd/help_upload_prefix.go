package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_upload_prefixCmd = &cobra.Command{
	Use:   "prefix",
	Short: "Options for uploading to a prefix.dev server. Authentication is used from the keychain / auth-file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_upload_prefixCmd).Standalone()

	help_uploadCmd.AddCommand(help_upload_prefixCmd)
}
