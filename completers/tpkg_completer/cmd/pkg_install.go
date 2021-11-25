package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/tpkg_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var pkg_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs a package in the current project, or downloads all dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_installCmd).Standalone()
	pkg_installCmd.Flags().Bool("local", false, "Treat package argument as local path")
	pkg_installCmd.Flags().String("name", "", "The name used for the 'import' clause")
	pkg_installCmd.Flags().Bool("recompute", false, "Recompute dependencies")
	pkgCmd.AddCommand(pkg_installCmd)

	carapace.Gen(pkg_installCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if pkg_installCmd.Flag("local").Changed {
				return carapace.ActionDirectories()
			}

			return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					if strings.HasPrefix(c.CallbackValue, "github.com/") {
						c.CallbackValue = strings.TrimPrefix(c.CallbackValue, "github.com/")
						return gh.ActionOwnerRepositories().Invoke(c).Prefix("github.com/").ToA()
					}

					return carapace.Batch(
						action.ActionPackages(),
						carapace.ActionValues("github.com/"),
					).ToA()
				case 1:
					return action.ActionPackageVersions(c.Parts[0])
				default:
					return carapace.ActionValues()
				}
			})
		}),
	)
}
