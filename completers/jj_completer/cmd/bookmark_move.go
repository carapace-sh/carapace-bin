package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var bookmark_moveCmd = &cobra.Command{
	Use:     "move [OPTIONS] <--from <REVISIONS>|NAME>",
	Short:   "Move existing bookmarks to target revision",
	Aliases: []string{"s"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmark_moveCmd).Standalone()

	bookmark_moveCmd.Flags().BoolP("allow-backwards", "B", false, "Allow moving the bookmark backwards or sideways")
	bookmark_moveCmd.Flags().String("from", "@", "Move part of this change into the destination")
	bookmark_moveCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bookmark_moveCmd.Flags().String("to", "@", "Move part of the source into this change")
	bookmarkCmd.AddCommand(bookmark_moveCmd)

	carapace.Gen(bookmark_moveCmd).FlagCompletion(carapace.ActionMap{
		"from": jj.ActionRevs(jj.RevOption{}.Default()),
		"to":   jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(bookmark_moveCmd).PositionalAnyCompletion(
		jj.ActionLocalBookmarks().FilterArgs(),
	)
}
