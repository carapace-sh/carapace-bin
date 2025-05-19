package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var stashCmd = &cobra.Command{
	Use:     "stash",
	Short:   "Stash the changes in a dirty working directory away",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(stashCmd).Standalone()
	rootCmd.AddCommand(stashCmd)

	carapace.Gen(stashCmd).DashAnyCompletion(
		git.ActionChanges(git.ChangeOpts{Unstaged: true}).FilterArgs(),
	)
}
