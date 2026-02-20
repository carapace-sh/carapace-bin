package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_quetzCmd = &cobra.Command{
	Use:   "quetz",
	Short: "Upload to a Quetz server. Authentication is used from the keychain / auth-file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_quetzCmd).Standalone()

	upload_quetzCmd.Flags().StringP("api-key", "a", "", "The Quetz API key, if none is provided, the token is read from the keychain / auth-file")
	upload_quetzCmd.Flags().StringP("channel", "c", "", "The URL to your channel")
	upload_quetzCmd.Flags().StringP("url", "u", "", "The URL to your Quetz server")
	upload_quetzCmd.MarkFlagRequired("channel")
	upload_quetzCmd.MarkFlagRequired("url")
	uploadCmd.AddCommand(upload_quetzCmd)
}
