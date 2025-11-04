package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var httpPushCmd = &cobra.Command{
	Use:     "http-push",
	Short:   "Push objects over HTTP/DAV to another repository",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_synching].ID,
}

func init() {
	carapace.Gen(httpPushCmd).Standalone()

	httpPushCmd.Flags().BoolS("D", "D", false, "remove <ref> from remote repository")
	httpPushCmd.Flags().Bool("all", false, "verify all objects in the entire local ref’s history exist")
	httpPushCmd.Flags().BoolS("d", "d", false, "remove <ref> from remote repository")
	httpPushCmd.Flags().Bool("dry-run", false, "do everything except actually send the updates")
	httpPushCmd.Flags().Bool("force", false, "disable ancestor check")
	httpPushCmd.Flags().Bool("verbose", false, "report the list of objects being walked and successfully sent")
	rootCmd.AddCommand(httpPushCmd)

	httpPushCmd.MarkFlagsMutuallyExclusive("D", "d")

	// TODO positional completion
}
