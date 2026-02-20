package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_upload_s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Options for uploading to S3",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_upload_s3Cmd).Standalone()

	help_uploadCmd.AddCommand(help_upload_s3Cmd)
}
