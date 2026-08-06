package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var tag_trackCmd = &cobra.Command{
	Use:     "track [OPTIONS] <NAMES>...",
	Short:   "Start tracking given remote tags",
	Aliases: []string{"t"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tag_trackCmd).Standalone()

	tag_trackCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	tag_trackCmd.Flags().StringSlice("remote", nil, "Remote names to track")
	tagCmd.AddCommand(tag_trackCmd)

	carapace.Gen(tag_trackCmd).FlagCompletion(carapace.ActionMap{
		"remote": jj.ActionRemotes(),
	})

	carapace.Gen(tag_trackCmd).PositionalAnyCompletion(
		jj.ActionTags().FilterArgs(),
	)
}
