package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/git"
	"github.com/spf13/cobra"
)

var stash_showCmd = &cobra.Command{
	Use:   "show",
	Short: "show the changes recorded in the stash entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_showCmd).Standalone()

	stashCmd.AddCommand(stash_showCmd)

	carapace.Gen(stash_showCmd).PositionalCompletion(git.ActionStashes())
}
