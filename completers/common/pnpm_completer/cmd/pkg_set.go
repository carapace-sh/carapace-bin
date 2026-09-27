package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pkg_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a value in package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_setCmd).Standalone()

	pkg_setCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pkgCmd.AddCommand(pkg_setCmd)
}
