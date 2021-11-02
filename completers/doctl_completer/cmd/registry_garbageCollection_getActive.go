package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_garbageCollection_getActiveCmd = &cobra.Command{
	Use:   "get-active",
	Short: "Retrieve information about the currently-active garbage collection for a container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_garbageCollection_getActiveCmd).Standalone()
	registry_garbageCollection_getActiveCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `UUID`, `RegistryName`, `Status`, `CreatedAt`, `UpdatedAt`, `BlobsDeleted`, `FreedBytes`")
	registry_garbageCollection_getActiveCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	registry_garbageCollectionCmd.AddCommand(registry_garbageCollection_getActiveCmd)
}
