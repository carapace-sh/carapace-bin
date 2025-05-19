package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var repo_publishCmd = &cobra.Command{
	Use:   "publish <command> [flags]",
	Short: "Publishes resources in the project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_publishCmd).Standalone()

	repoCmd.AddCommand(repo_publishCmd)
}
