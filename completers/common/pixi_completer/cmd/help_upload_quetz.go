package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_upload_quetzCmd = &cobra.Command{
	Use:   "quetz",
	Short: "Upload to a Quetz server. Authentication is used from the keychain / auth-file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_upload_quetzCmd).Standalone()

	help_uploadCmd.AddCommand(help_upload_quetzCmd)
}
