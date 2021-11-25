package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pkg_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a new package and lock file in the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_initCmd).Standalone()
	pkgCmd.AddCommand(pkg_initCmd)
}
