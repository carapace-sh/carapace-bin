package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Upload distributions to an index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(publishCmd).Standalone()

	publishCmd.Flags().String("check-url", "", "Check an index URL for existing files to skip duplicate uploads")
	publishCmd.Flags().Bool("dry-run", false, "Perform a dry run without uploading files")
	publishCmd.Flags().String("index", "", "The name of an index in the configuration to use for publishing.")
	publishCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for remote requirements files")
	publishCmd.Flags().Bool("no-attestations", false, "Do not upload attestations for the published files")
	publishCmd.Flags().StringP("password", "p", "", "The password for the upload")
	publishCmd.Flags().String("publish-url", "", "The URL of the upload endpoint (not the index URL)")
	publishCmd.Flags().Bool("skip-existing", false, "")
	publishCmd.Flags().StringP("token", "t", "", "The token for the upload")
	publishCmd.Flags().String("trusted-publishing", "", "Configure trusted publishing")
	publishCmd.Flags().StringP("username", "u", "", "The username for the upload")
	publishCmd.Flag("skip-existing").Hidden = true
	rootCmd.AddCommand(publishCmd)
	carapace.Gen(publishCmd).FlagCompletion(carapace.ActionMap{
		"keyring-provider": uv.ActionKeyringProviders(),
		"trusted-publishing": carapace.ActionValuesDescribed(
			"automatic", "Attempt trusted publishing when we're in a supported environment, continue if that fails",
			"always", "",
			"never", "",
		),
	})
	carapace.Gen(publishCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
