package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var lastModifiedCmd = &cobra.Command{
	Use:     "last-modified",
	Short:   "EXPERIMENTAL: Show when files were last modified",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_interrogator].ID,
}

func init() {
	carapace.Gen(lastModifiedCmd).Standalone()

	lastModifiedCmd.Flags().String("max-depth", "", "maximum tree depth to recurse")
	lastModifiedCmd.Flags().BoolP("recursive", "r", false, "recurse into subtrees")
	lastModifiedCmd.Flags().BoolP("show-trees", "t", false, "show tree entries when recursing into subtrees")
	lastModifiedCmd.Flags().BoolS("z", "z", false, "lines are separated with NUL character")
	rootCmd.AddCommand(lastModifiedCmd)

	carapace.Gen(lastModifiedCmd).PositionalAnyCompletion(
		carapace.Batch(
			git.ActionRefRanges(git.RefOption{}.Default()),
			carapace.ActionFiles(),
		).ToA(),
	)

	carapace.Gen(lastModifiedCmd).DashAnyCompletion(
		carapace.ActionPositional(lastModifiedCmd),
	)
}
