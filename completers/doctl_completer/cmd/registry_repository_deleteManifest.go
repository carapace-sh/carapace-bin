package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_repository_deleteManifestCmd = &cobra.Command{
	Use:   "delete-manifest",
	Short: "Delete one or more container repository manifests by digest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_repository_deleteManifestCmd).Standalone()
	registry_repository_deleteManifestCmd.Flags().BoolP("force", "f", false, "Force manifest deletion")
	registry_repositoryCmd.AddCommand(registry_repository_deleteManifestCmd)
}
