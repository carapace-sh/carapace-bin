package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var bookmark_setCmd = &cobra.Command{
	Use:     "set [OPTIONS] <NAMES>...",
	Short:   "Update a given bookmark to point to a certain commit",
	Aliases: []string{"s"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bookmark_setCmd).Standalone()

	bookmark_setCmd.Flags().BoolP("allow-backwards", "B", false, "Allow moving the bookmark backwards or sideways")
	bookmark_setCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	bookmark_setCmd.Flags().StringP("revision", "r", "", "The bookmark's target revision")
	bookmarkCmd.AddCommand(bookmark_setCmd)

	carapace.Gen(bookmark_setCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(bookmark_setCmd).PositionalAnyCompletion(
		jj.ActionLocalBookmarks().FilterArgs(),
	)
}
