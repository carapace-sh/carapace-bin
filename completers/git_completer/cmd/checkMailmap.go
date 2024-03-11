package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkMailmapCmd = &cobra.Command{
	Use:     "check-mailmap",
	Short:   "Show canonical names and email addresses of contacts",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(checkMailmapCmd).Standalone()

	checkMailmapCmd.Flags().Bool("stdin", false, "also read contacts from stdin")
	rootCmd.AddCommand(checkMailmapCmd)
}
