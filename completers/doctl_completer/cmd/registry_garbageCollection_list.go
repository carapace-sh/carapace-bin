package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_garbageCollection_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Retrieve information about past garbage collections for a container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_garbageCollection_listCmd).Standalone()
	registry_garbageCollection_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `UUID`, `RegistryName`, `Status`, `CreatedAt`, `UpdatedAt`, `BlobsDeleted`, `FreedBytes`")
	registry_garbageCollection_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	registry_garbageCollectionCmd.AddCommand(registry_garbageCollection_listCmd)
}
