package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var pkg_registry_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkg_registry_addCmd).Standalone()
	pkg_registry_addCmd.Flags().Bool("local", false, "Registry is local")
	pkg_registryCmd.AddCommand(pkg_registry_addCmd)

	carapace.Gen(pkg_registry_addCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if pkg_registry_addCmd.Flag("local").Changed {
				return carapace.ActionDirectories()
			}
			return git.ActionRepositorySearch()
		}),
	)
}
