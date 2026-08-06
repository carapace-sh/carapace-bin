package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var tag_untrackCmd = &cobra.Command{
	Use:     "untrack [OPTIONS] <NAMES>...",
	Short:   "Stop tracking given remote tags",
	Aliases: []string{"u"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tag_untrackCmd).Standalone()

	tag_untrackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	tag_untrackCmd.Flags().StringSlice("remote", nil, "Remote names to untrack")
	tagCmd.AddCommand(tag_untrackCmd)

	carapace.Gen(tag_untrackCmd).FlagCompletion(carapace.ActionMap{
		"remote": jj.ActionRemotes(),
	})

	carapace.Gen(tag_untrackCmd).PositionalAnyCompletion(
		jj.ActionTags().FilterArgs(),
	)
}
