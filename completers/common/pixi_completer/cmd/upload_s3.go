package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upload_s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Upload to S3",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upload_s3Cmd).Standalone()

	upload_s3Cmd.Flags().String("access-key-id", "", "The S3 access key ID")
	upload_s3Cmd.Flags().String("addressing-style", "", "The S3 addressing style")
	upload_s3Cmd.Flags().StringP("channel", "c", "", "The channel to upload the package to")
	upload_s3Cmd.Flags().String("endpoint-url", "", "The S3 endpoint URL")
	upload_s3Cmd.Flags().Bool("force", false, "Force upload even if the package already exists")
	upload_s3Cmd.Flags().String("force-path-style", "", "Force path-style addressing")
	upload_s3Cmd.Flags().String("region", "", "The S3 region")
	upload_s3Cmd.Flags().String("secret-access-key", "", "The S3 secret access key")
	upload_s3Cmd.Flags().String("session-token", "", "The S3 session token")
	uploadCmd.AddCommand(upload_s3Cmd)

	carapace.Gen(upload_s3Cmd).FlagCompletion(carapace.ActionMap{
		"addressing-style": carapace.ActionValues("virtual-host", "path"),
		"force-path-style": carapace.ActionValues("true", "false"),
	})
}
