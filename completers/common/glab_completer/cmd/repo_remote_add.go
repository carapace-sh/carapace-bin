package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_remote_addCmd = &cobra.Command{
	Use:   "add <namespace/project>",
	Short: "Add a Git remote for a GitLab project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_remote_addCmd).Standalone()

	repo_remote_addCmd.Flags().StringP("name", "n", "", "Name for the remote (default: first path component)")
	repo_remote_addCmd.Flags().StringP("protocol", "p", "", "Git protocol: ssh, https (default: git_protocol config)")
	repo_remoteCmd.AddCommand(repo_remote_addCmd)
}
