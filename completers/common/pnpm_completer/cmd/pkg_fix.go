package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pkg_fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "Auto corrects common errors in package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_fixCmd).Standalone()

	pkg_fixCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pkgCmd.AddCommand(pkg_fixCmd)
}
