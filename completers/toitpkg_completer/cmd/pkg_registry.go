package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pkg_registryCmd = &cobra.Command{
	Use:   "registry",
	Short: "Manages registries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_registryCmd).Standalone()
	pkgCmd.AddCommand(pkg_registryCmd)
}
