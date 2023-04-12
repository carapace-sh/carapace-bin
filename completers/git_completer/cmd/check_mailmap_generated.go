package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var check_mailmapCmd = &cobra.Command{
	Use:     "check-mailmap",
	Short:   "Show canonical names and email addresses of contacts",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(check_mailmapCmd).Standalone()
	check_mailmapCmd.Flags().Bool("stdin", false, "also read contacts from stdin")
	rootCmd.AddCommand(check_mailmapCmd)
}
