package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/toitpkg_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pkg_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls the package with the given name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_uninstallCmd).Standalone()
	pkgCmd.AddCommand(pkg_uninstallCmd)

	carapace.Gen(pkg_uninstallCmd).PositionalCompletion(
		action.ActionDependencies(pkg_uninstallCmd),
	)
}
