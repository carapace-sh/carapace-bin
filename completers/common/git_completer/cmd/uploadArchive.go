package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uploadArchiveCmd = &cobra.Command{
	Use:     "upload-archive",
	Short:   "Send archive back to git-archive",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_synching].ID,
}

func init() {
	carapace.Gen(uploadArchiveCmd).Standalone()

	rootCmd.AddCommand(uploadArchiveCmd)

	carapace.Gen(uploadArchiveCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
