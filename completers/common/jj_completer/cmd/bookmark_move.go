package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var bookmark_moveCmd = &cobra.Command{
	Use:     "move [OPTIONS] <--from <REVISIONS>|NAME>",
	Short:   "Move existing bookmarks to target revision",
	Aliases: []string{"m"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmark_moveCmd).Standalone()

	bookmark_moveCmd.Flags().BoolP("allow-backwards", "B", false, "Allow moving bookmarks backwards or sideways")
	bookmark_moveCmd.Flags().StringSliceP("from", "f", nil, "Move bookmarks from the given revisions")
	bookmark_moveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bookmark_moveCmd.Flags().StringP("to", "t", "@", "Move bookmarks to this revision")
	bookmarkCmd.AddCommand(bookmark_moveCmd)

	carapace.Gen(bookmark_moveCmd).FlagCompletion(carapace.ActionMap{
		"from": jj.ActionRevs(jj.RevOption{}.Default()),
		"to":   jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(bookmark_moveCmd).PositionalAnyCompletion(
		jj.ActionLocalBookmarks().FilterArgs(),
	)
}
