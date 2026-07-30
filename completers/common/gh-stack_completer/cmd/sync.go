package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:     "sync",
	Short:   "Sync the current stack with the remote",
	GroupID: "remote",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syncCmd).Standalone()

	syncCmd.Flags().Bool("prune", false, "Delete local branches for merged PRs")
	syncCmd.Flags().String("remote", "", "Remote to fetch from and push to (defaults to auto-detected remote)")
	rootCmd.AddCommand(syncCmd)

	carapace.Gen(syncCmd).FlagCompletion(carapace.ActionMap{
		"remote": git.ActionRemotes(),
	})
}
