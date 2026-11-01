package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_infoCmd).Standalone()

	help_repositoryCmd.AddCommand(help_repository_infoCmd)
}
