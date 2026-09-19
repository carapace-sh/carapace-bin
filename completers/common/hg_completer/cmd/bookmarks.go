package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/hg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var bookmarksCmd = &cobra.Command{
	Use:     "bookmarks",
	Short:   "create a new bookmark or list existing bookmarks",
	Aliases: []string{"bookmark"},
	GroupID: groups[group_change_organization].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmarksCmd).Standalone()

	bookmarksCmd.Flags().BoolP("delete", "d", false, "delete a given bookmark")
	bookmarksCmd.Flags().BoolP("force", "f", false, "force")
	bookmarksCmd.Flags().BoolP("inactive", "i", false, "mark a bookmark inactive")
	bookmarksCmd.Flags().BoolP("list", "l", false, "list existing bookmarks")
	bookmarksCmd.Flags().StringP("rename", "m", "", "rename a given bookmark")
	bookmarksCmd.Flags().StringP("rev", "r", "", "revision for bookmark action")
	bookmarksCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(bookmarksCmd)

	carapace.Gen(bookmarksCmd).FlagCompletion(carapace.ActionMap{
		"rename": action.ActionBookmarks(),
		"rev":    action.ActionRevisions(),
	})

	carapace.Gen(bookmarksCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if bookmarksCmd.Flags().Changed("delete") || bookmarksCmd.Flags().Changed("rename") {
				return action.ActionBookmarks()
			}
			return carapace.ActionValues()
		}),
	)
}
