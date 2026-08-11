package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_listCmd).Standalone()

	help_repositoryCmd.AddCommand(help_repository_listCmd)
}
