package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pkg_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves a value from package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_getCmd).Standalone()

	pkg_getCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pkgCmd.AddCommand(pkg_getCmd)
}
