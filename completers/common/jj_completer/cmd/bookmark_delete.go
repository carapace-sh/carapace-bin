package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var bookmark_deleteCmd = &cobra.Command{
	Use:     "delete [OPTIONS] [NAMES]...",
	Short:   "Delete an existing bookmark and propagate the deletion to remotes on the next push",
	Aliases: []string{"d"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmark_deleteCmd).Standalone()

	bookmark_deleteCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bookmarkCmd.AddCommand(bookmark_deleteCmd)

	carapace.Gen(bookmark_deleteCmd).PositionalAnyCompletion(
		jj.ActionLocalBookmarks().FilterArgs(),
	)
}
