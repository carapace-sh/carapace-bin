package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var bookmark_listCmd = &cobra.Command{
	Use:     "list [OPTIONS] [NAMES]...",
	Short:   "List bookmarks and their targets",
	Aliases: []string{"l"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmark_listCmd).Standalone()

	bookmark_listCmd.Flags().BoolP("all-remotes", "a", false, "Show all tracking and non-tracking remote bookmarks including the ones whose targets are synchronized with the local bookmarks")
	bookmark_listCmd.Flags().BoolP("conflicted", "c", false, "Show only conflicted bookmarks")
	bookmark_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bookmark_listCmd.Flags().StringSlice("remote", nil, "Show all tracking and non-tracking remote bookmarks belonging to this remote")
	bookmark_listCmd.Flags().StringSliceP("revisions", "r", nil, "Show bookmarks whose local targets are in the given revisions")
	bookmark_listCmd.Flags().StringP("template", "T", "", "Render each bookmark using the given template")

	bookmark_listCmd.MarkFlagsMutuallyExclusive("all-remotes", "conflicted")

	bookmarkCmd.AddCommand(bookmark_listCmd)

	carapace.Gen(bookmark_listCmd).FlagCompletion(carapace.ActionMap{
		"remote":    jj.ActionRemotes(),
		"revisions": jj.ActionRevSets(jj.RevOption{}.Default()).UniqueList(","),
	})

	carapace.Gen(bookmark_listCmd).PositionalAnyCompletion(
		jj.ActionLocalBookmarks().FilterArgs(),
	)
}
