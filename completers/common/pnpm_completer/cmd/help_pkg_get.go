package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_pkg_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves a value from package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_pkg_getCmd).Standalone()

	help_pkgCmd.AddCommand(help_pkg_getCmd)
}
