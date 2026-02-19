package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_artifactoryCmd = &cobra.Command{
	Use:   "artifactory",
	Short: "Upload to an Artifactory channel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_artifactoryCmd).Standalone()

	upload_artifactoryCmd.Flags().StringP("channel", "c", "", "The channel to upload the package to")
	upload_artifactoryCmd.Flags().String("password", "", "The password for authentication")
	upload_artifactoryCmd.Flags().StringP("token", "t", "", "The token for authentication")
	upload_artifactoryCmd.Flags().StringP("url", "u", "", "The URL to the Artifactory server")
	upload_artifactoryCmd.Flags().String("username", "", "The username for authentication")
	uploadCmd.AddCommand(upload_artifactoryCmd)
}
