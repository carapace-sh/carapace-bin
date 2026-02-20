package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_help_s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Options for uploading to S3",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_help_s3Cmd).Standalone()

	upload_helpCmd.AddCommand(upload_help_s3Cmd)
}
