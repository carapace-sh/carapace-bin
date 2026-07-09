package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:     "switch",
	Short:   "Switch to another branch",
	Aliases: []string{"checkout", "co"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(switchCmd).Standalone()

	switchCmd.Flags().BoolP("create", "c", false, "Create the branch first, then switch to it")
	switchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(switchCmd)

	carapace.Gen(switchCmd).PositionalCompletion(
		carapace.Batch(
			git.ActionRemoteBranchNames(""),
			git.ActionRefs(git.RefOption{LocalBranches: true}),
		).ToA(),
	)

	carapace.Gen(switchCmd).DashAnyCompletion(
		carapace.ActionPositional(switchCmd),
	)
}
