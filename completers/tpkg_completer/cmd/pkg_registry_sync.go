package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pkg_registry_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronizes all registries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_registry_syncCmd).Standalone()
	pkg_registryCmd.AddCommand(pkg_registry_syncCmd)
}
