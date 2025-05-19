package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var bookmark_createCmd = &cobra.Command{
	Use:     "create [OPTIONS] <NAMES>...",
	Short:   "Create a new bookmark",
	Aliases: []string{"c"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmark_createCmd).Standalone()

	bookmark_createCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bookmark_createCmd.Flags().StringP("revision", "r", "", "The bookmark's target revision")
	bookmarkCmd.AddCommand(bookmark_createCmd)

	carapace.Gen(bookmark_createCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(bookmark_createCmd).PositionalAnyCompletion(
		jj.ActionLocalBookmarks().FilterArgs(),
	)
}
