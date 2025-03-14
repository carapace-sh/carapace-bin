package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var bookmark_forgetCmd = &cobra.Command{
	Use:     "forget [OPTIONS] [NAMES]...",
	Short:   "Forget everything about a bookmark, including its local and remote targets",
	Aliases: []string{"f"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmark_forgetCmd).Standalone()

	bookmark_forgetCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bookmark_forgetCmd.Flags().Bool("include-remotes", false, "When forgetting a local bookmark, also forget any corresponding remote bookmarks")
	bookmarkCmd.AddCommand(bookmark_forgetCmd)

	carapace.Gen(bookmark_forgetCmd).PositionalAnyCompletion(
		jj.ActionLocalBookmarks().FilterArgs(),
	)
}
