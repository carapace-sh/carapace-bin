package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/paru_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/paru"
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
	common.AddNewFlags(getpkgbuildCmd)

	carapace.Gen(getpkgbuildCmd).PositionalAnyCompletion(
		paru.ActionPackageSearch().FilterArgs(),
	)
}
