package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push all branches in the current stack to the remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	pushCmd.Flags().String("remote", "", "Remote to push to (defaults to auto-detected remote)")
	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).FlagCompletion(carapace.ActionMap{
		"remote": git.ActionRemotes(),
	})
}
