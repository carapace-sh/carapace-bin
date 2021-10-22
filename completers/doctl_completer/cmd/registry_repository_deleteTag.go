package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_repository_deleteTagCmd = &cobra.Command{
	Use:   "delete-tag",
	Short: "Delete one or more container repository tags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_repository_deleteTagCmd).Standalone()
	registry_repository_deleteTagCmd.Flags().BoolP("force", "f", false, "Force tag deletion")
	registry_repositoryCmd.AddCommand(registry_repository_deleteTagCmd)
}
