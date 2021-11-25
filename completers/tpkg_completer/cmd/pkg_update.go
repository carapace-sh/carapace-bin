package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pkg_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates all packages to their newest versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_updateCmd).Standalone()
	pkgCmd.AddCommand(pkg_updateCmd)
}
