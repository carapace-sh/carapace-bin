package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_artifactoryCmd = &cobra.Command{
	Use:   "artifactory",
	Short: "Options for uploading to a Artifactory channel. Authentication is used from the keychain / auth-file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_artifactoryCmd).Standalone()

	upload_artifactoryCmd.Flags().StringP("channel", "c", "", "The URL to your channel")
	upload_artifactoryCmd.Flags().StringP("token", "t", "", "Your Artifactory token")
	upload_artifactoryCmd.Flags().StringP("url", "u", "", "The URL to your Artifactory server")
	upload_artifactoryCmd.MarkFlagRequired("channel")
	upload_artifactoryCmd.MarkFlagRequired("url")
	uploadCmd.AddCommand(upload_artifactoryCmd)
}
