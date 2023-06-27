package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/pacman"
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

	getpkgbuildCmd.Flags().BoolP("-f", "--force", false, "Force download for existing ABS packages")
	getpkgbuildCmd.Flags().BoolP("-p", "--print", false, "Print pkgbuild of packages")

	carapace.Gen(getpkgbuildCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return pacman.ActionPackageSearch().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
