package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var release_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List releases in a repository",
	GroupID: "General commands",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_listCmd).Standalone()

	release_listCmd.Flags().Bool("exclude-drafts", false, "Exclude draft releases")
	release_listCmd.Flags().Bool("exclude-pre-releases", false, "Exclude pre-releases")
	release_listCmd.Flags().StringP("limit", "L", "", "Maximum number of items to fetch")
	releaseCmd.AddCommand(release_listCmd)
}
