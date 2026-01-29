package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uploadPackCmd = &cobra.Command{
	Use:     "upload-pack",
	Short:   "Send objects packed back to git-fetch-pack",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_synching].ID,
}

func init() {
	carapace.Gen(uploadPackCmd).Standalone()

	uploadPackCmd.Flags().Bool("http-backend-info-refs", false, "Used by git-http-backend(1) to serve up $GIT_URL/info/refs?service=git-upload-pack requests")
	uploadPackCmd.Flags().Bool("stateless-rpc", false, "Perform only a single read-write cycle with stdin and stdout")
	uploadPackCmd.Flags().Bool("strict", false, "Do not try <directory>/.git/ if <directory> is not a Git directory")
	uploadPackCmd.Flags().String("timeout", "", "Interrupt transfer after <n> seconds of inactivity")
	rootCmd.AddCommand(uploadPackCmd)

	carapace.Gen(uploadPackCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
