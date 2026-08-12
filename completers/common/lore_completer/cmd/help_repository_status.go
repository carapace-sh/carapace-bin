package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show current repository status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_statusCmd).Standalone()

	help_repositoryCmd.AddCommand(help_repository_statusCmd)
}
