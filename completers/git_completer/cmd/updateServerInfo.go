package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateServerInfoCmd = &cobra.Command{
	Use:     "update-server-info",
	Short:   "Update auxiliary info file to help dumb servers",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_synching].ID,
}

func init() {
	carapace.Gen(updateServerInfoCmd).Standalone()

	updateServerInfoCmd.Flags().BoolP("force", "f", false, "update the info files from scratch")
	rootCmd.AddCommand(updateServerInfoCmd)
}
