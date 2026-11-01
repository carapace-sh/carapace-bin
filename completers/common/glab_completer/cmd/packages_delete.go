package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var packages_deleteCmd = &cobra.Command{
	Use:     "delete [<id> | --id <id> | --name <name>] [flags]",
	Short:   "Delete a package from a project's package registry.",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packages_deleteCmd).Standalone()

	packages_deleteCmd.Flags().String("id", "", "ID of the package to remove.")
	packages_deleteCmd.Flags().String("name", "", "Name of the package to remove.")
	packages_deleteCmd.Flags().String("package-type", "generic", "Type of the package to remove, used with --name. One of: cargo, composer, conan, debian, generic, golang, helm, maven, ml_model, npm, nuget, pypi, rpm, rubygems, terraform_module.")
	packages_deleteCmd.Flags().String("version", "", "Version of the package.")
	packages_deleteCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt.")
	packagesCmd.AddCommand(packages_deleteCmd)

	carapace.Gen(packages_deleteCmd).FlagCompletion(carapace.ActionMap{
		"id":           action.ActionPackages(packages_deleteCmd),
		"name":         action.ActionPackageNames(packages_deleteCmd),
		"package-type": carapace.ActionValues("cargo", "composer", "conan", "debian", "generic", "golang", "helm", "maven", "ml_model", "npm", "nuget", "pypi", "rpm", "rubygems", "terraform_module"),
	})
	carapace.Gen(packages_deleteCmd).PositionalCompletion(
		action.ActionPackages(packages_deleteCmd),
	)
}
