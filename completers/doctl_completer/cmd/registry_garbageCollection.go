package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_garbageCollectionCmd = &cobra.Command{
	Use:   "garbage-collection",
	Short: "Display commands for garbage collection for a container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_garbageCollectionCmd).Standalone()
	registryCmd.AddCommand(registry_garbageCollectionCmd)
}
