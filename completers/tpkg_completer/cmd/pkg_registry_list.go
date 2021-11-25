package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pkg_registry_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List registries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_registry_listCmd).Standalone()
	pkg_registryCmd.AddCommand(pkg_registry_listCmd)
}
