package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_help_prefixCmd = &cobra.Command{
	Use:   "prefix",
	Short: "Options for uploading to a prefix.dev server. Authentication is used from the keychain / auth-file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_help_prefixCmd).Standalone()

	upload_helpCmd.AddCommand(upload_help_prefixCmd)
}
