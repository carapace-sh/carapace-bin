package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_membersCmd = &cobra.Command{
	Use:   "members <command> [flags]",
	Short: "Manage project members.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_membersCmd).Standalone()

	repoCmd.AddCommand(repo_membersCmd)
}
