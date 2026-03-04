package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var bookmark_advanceCmd = &cobra.Command{
	Use:     "advance",
	Short:   "Advance the closest bookmarks to a target revision",
	Aliases: []string{"a"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmark_advanceCmd).Standalone()

	bookmark_advanceCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bookmark_advanceCmd.Flags().StringP("to", "t", "", "Move bookmarks to this revision")
	bookmarkCmd.AddCommand(bookmark_advanceCmd)

	carapace.Gen(bookmark_advanceCmd).FlagCompletion(carapace.ActionMap{
		"to": jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(bookmark_advanceCmd).PositionalAnyCompletion(
		jj.ActionLocalBookmarks(),
	)
}
