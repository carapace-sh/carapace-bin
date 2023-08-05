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

	getpkgbuildCmd.Flags().BoolP("comments", "c", false, "Print AUR comments for pkgbuild")
	getpkgbuildCmd.Flags().BoolP("print", "p", false, "Print pkgbuild to stdout")
	getpkgbuildCmd.Flags().BoolP("ssh", "s", false, "Clone package using SSH")

	carapace.Gen(getpkgbuildCmd).PositionalAnyCompletion(
		pacman.ActionPackageSearch().FilterArgs(),
	)
}
