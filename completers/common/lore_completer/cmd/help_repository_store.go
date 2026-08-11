package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_repository_storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Access the repository data store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_repository_storeCmd).Standalone()

	help_repositoryCmd.AddCommand(help_repository_storeCmd)
}
