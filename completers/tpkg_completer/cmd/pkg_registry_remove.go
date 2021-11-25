package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/tpkg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pkg_registry_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_registry_removeCmd).Standalone()
	pkg_registryCmd.AddCommand(pkg_registry_removeCmd)

	carapace.Gen(pkg_registry_removeCmd).PositionalCompletion(
		action.ActionRegistries(),
	)
}
