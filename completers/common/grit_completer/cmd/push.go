package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Publish the current branch to its remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	pushCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pushCmd.Flags().BoolP("tags", "t", false, "Push tags instead of the current branch. Pushes every local tag under `refs/tags/` to the remote")
	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).PositionalCompletion(
		git.ActionRemotes(),
	)
}
