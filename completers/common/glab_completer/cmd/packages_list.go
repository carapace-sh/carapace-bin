package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jq/pkg/actions/tools/jq"
	"github.com/spf13/cobra"
)

var packages_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List packages in a project's package registry.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packages_listCmd).Standalone()

	packages_listCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	packages_listCmd.Flags().StringP("name", "n", "", "Filter packages by name (substring match).")
	packages_listCmd.Flags().String("package-type", "", "Filter packages by type. One of: composer, conan, debian, generic, golang, helm, maven, npm, nuget, pypi, terraform_module.")
	packages_listCmd.Flags().StringP("page", "p", "1", "Page number.")
	packages_listCmd.Flags().StringP("per-page", "P", "30", "Number of items to list per page.")
	packagesCmd.AddCommand(packages_listCmd)

	carapace.Gen(packages_listCmd).FlagCompletion(carapace.ActionMap{
		"jq":           jq.ActionFilters(),
		"package-type": carapace.ActionValues("composer", "conan", "debian", "generic", "golang", "helm", "maven", "npm", "nuget", "pypi", "terraform_module"),
	})
}
