package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var containerRegistry_repository_deleteCmd = &cobra.Command{
	Use:     "delete <repository-id> [flags]",
	Short:   "Delete a container registry repository.",
	Aliases: []string{"del"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(containerRegistry_repository_deleteCmd).Standalone()

	containerRegistry_repository_deleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt.")
	containerRegistry_repositoryCmd.AddCommand(containerRegistry_repository_deleteCmd)
}
