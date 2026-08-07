package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_remoteCmd = &cobra.Command{
	Use:   "remote <subcommand>",
	Short: "Manage Git remotes for a GitLab project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_remoteCmd).Standalone()

	repoCmd.AddCommand(repo_remoteCmd)
}
