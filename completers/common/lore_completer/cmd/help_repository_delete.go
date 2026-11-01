package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_deleteCmd).Standalone()

	help_repositoryCmd.AddCommand(help_repository_deleteCmd)
}
