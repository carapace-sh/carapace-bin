package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var bookmark_trackCmd = &cobra.Command{
	Use:   "track [OPTIONS] <NAMES>...",
	Short: "Start tracking given remote bookmarks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmark_trackCmd).Standalone()

	bookmark_trackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bookmarkCmd.AddCommand(bookmark_trackCmd)

	carapace.Gen(bookmark_trackCmd).PositionalAnyCompletion(
		jj.ActionRemoteBookmarks("").FilterArgs(),
	)
}
