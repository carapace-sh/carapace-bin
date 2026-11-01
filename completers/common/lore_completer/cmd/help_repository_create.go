package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a repository in the given directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_createCmd).Standalone()

	help_repositoryCmd.AddCommand(help_repository_createCmd)
}
