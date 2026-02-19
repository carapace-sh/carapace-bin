package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_quetzCmd = &cobra.Command{
	Use:   "quetz",
	Short: "Upload to a Quetz server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_quetzCmd).Standalone()

	upload_quetzCmd.Flags().StringP("api-key", "a", "", "The Quetz API key")
	upload_quetzCmd.Flags().StringP("channel", "c", "", "The channel to upload the package to")
	upload_quetzCmd.Flags().StringP("url", "u", "", "The URL to the Quetz server")
	uploadCmd.AddCommand(upload_quetzCmd)
}
