package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var remote_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Fetch updates for remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(remote_updateCmd).Standalone()

	remote_updateCmd.Flags().BoolP("prune", "p", false, "prune remotes after fetching")
	remoteCmd.AddCommand(remote_updateCmd)

	carapace.Gen(remote_updateCmd).PositionalAnyCompletion(
		git.ActionRemotes().FilterArgs(),
	)
}
