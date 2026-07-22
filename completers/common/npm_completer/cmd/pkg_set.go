package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var pkg_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value in package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_setCmd).Standalone()

	pkgCmd.AddCommand(pkg_setCmd)

	carapace.Gen(pkg_setCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return npm.ActionPackageJsonKeys().Invoke(c).Suffix("=").ToA()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
