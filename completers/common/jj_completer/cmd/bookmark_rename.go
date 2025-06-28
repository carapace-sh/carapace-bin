package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var bookmark_renameCmd = &cobra.Command{
	Use:     "rename",
	Short:   "Rename `old` bookmark name to `new` bookmark name",
	Aliases: []string{"r"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmark_renameCmd).Standalone()

	bookmark_renameCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bookmarkCmd.AddCommand(bookmark_renameCmd)

	carapace.Gen(bookmark_renameCmd).PositionalCompletion(
		jj.ActionLocalBookmarks(),
		jj.ActionLocalBookmarks(),
	)
}
