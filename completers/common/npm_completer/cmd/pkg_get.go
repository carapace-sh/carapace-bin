package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pkg_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a value from package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_getCmd).Standalone()

	pkgCmd.AddCommand(pkg_getCmd)
}
