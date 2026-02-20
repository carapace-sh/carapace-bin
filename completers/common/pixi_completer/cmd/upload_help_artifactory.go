package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_help_artifactoryCmd = &cobra.Command{
	Use:   "artifactory",
	Short: "Options for uploading to a Artifactory channel. Authentication is used from the keychain / auth-file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_help_artifactoryCmd).Standalone()

	upload_helpCmd.AddCommand(upload_help_artifactoryCmd)
}
