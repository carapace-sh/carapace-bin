package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_garbageCollection_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start garbage collection for a container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_garbageCollection_startCmd).Standalone()
	registry_garbageCollection_startCmd.Flags().Bool("exclude-unreferenced-blobs", false, "Exclude unreferenced blobs from garbage collection.")
	registry_garbageCollection_startCmd.Flags().BoolP("force", "f", false, "Run garbage collection without confirmation prompt")
	registry_garbageCollection_startCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `UUID`, `RegistryName`, `Status`, `CreatedAt`, `UpdatedAt`, `BlobsDeleted`, `FreedBytes`")
	registry_garbageCollection_startCmd.Flags().Bool("include-untagged-manifests", false, "Include untagged manifests in garbage collection.")
	registry_garbageCollection_startCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	registry_garbageCollectionCmd.AddCommand(registry_garbageCollection_startCmd)
}
