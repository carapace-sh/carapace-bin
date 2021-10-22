package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var registry_garbageCollection_cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel the currently-active garbage collection for a container registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registry_garbageCollection_cancelCmd).Standalone()
	registry_garbageCollectionCmd.AddCommand(registry_garbageCollection_cancelCmd)
}
