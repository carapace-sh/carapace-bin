package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yay"
	"github.com/spf13/cobra"
)

var getpkgbuildCmd = &cobra.Command{
	Use:     "getpkgbuild",
	Aliases: []string{"G"},
	Short:   "Get PKGBUILD from ABS or AUR",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getpkgbuildCmd).Standalone()

	getpkgbuildCmd.Flags().BoolP("force", "f", false, "Force download for existing ABS packages")
	getpkgbuildCmd.Flags().BoolP("print", "p", false, "Print pkgbuild of packages")

	carapace.Gen(getpkgbuildCmd).PositionalAnyCompletion(
		yay.ActionPackageSearch().FilterArgs(),
	)
}
