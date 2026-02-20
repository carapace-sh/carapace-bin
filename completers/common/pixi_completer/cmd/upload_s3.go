package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Options for uploading to S3",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_s3Cmd).Standalone()

	upload_s3Cmd.Flags().String("access-key-id", "", "The access key ID for the S3 bucket")
	upload_s3Cmd.Flags().String("addressing-style", "", "How to address the bucket")
	upload_s3Cmd.Flags().StringP("channel", "c", "", "The channel URL in the S3 bucket to upload the package to, e.g., `s3://my-bucket/my-channel`")
	upload_s3Cmd.Flags().String("endpoint-url", "", "The endpoint URL of the S3 backend")
	upload_s3Cmd.Flags().Bool("force", false, "Replace files if it already exists")
	upload_s3Cmd.Flags().String("region", "", "The region of the S3 backend")
	upload_s3Cmd.Flags().String("secret-access-key", "", "The secret access key for the S3 bucket")
	upload_s3Cmd.Flags().String("session-token", "", "The session token for the S3 bucket")
	upload_s3Cmd.MarkFlagRequired("channel")
	uploadCmd.AddCommand(upload_s3Cmd)

	carapace.Gen(upload_s3Cmd).FlagCompletion(carapace.ActionMap{
		"addressing-style": carapace.ActionValues("virtual-host", "path"),
	})
}
