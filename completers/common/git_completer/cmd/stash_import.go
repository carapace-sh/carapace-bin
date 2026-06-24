package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var stash_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import stash entries from a commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_importCmd).Standalone()

	stashCmd.AddCommand(stash_importCmd)

	carapace.Gen(stash_importCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
