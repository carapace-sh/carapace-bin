package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var bookmark_untrackCmd = &cobra.Command{
	Use:   "untrack",
	Short: "Stop tracking given remote bookmarks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmark_untrackCmd).Standalone()

	bookmark_untrackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bookmarkCmd.AddCommand(bookmark_untrackCmd)

	carapace.Gen(bookmark_untrackCmd).PositionalAnyCompletion(
		jj.ActionRemoteBookmarks("").FilterArgs(), // TODO tracked bookmarks
	)
}
