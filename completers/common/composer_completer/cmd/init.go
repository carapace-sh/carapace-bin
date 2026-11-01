package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a basic composer.json file in current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().String("author", "", "Author name of package")
	initCmd.Flags().StringP("autoload", "a", "", "Add PSR-4 autoload mapping. Maps your package's namespace to the provided directory. (Expects a relative path, e.g. src/)")
	initCmd.Flags().String("description", "", "Description of package")
	initCmd.Flags().String("homepage", "", "Homepage of package")
	initCmd.Flags().StringP("license", "l", "", "License of package")
	initCmd.Flags().String("name", "", "Name of the package")
	initCmd.Flags().StringArray("repository", nil, "Add custom repositories, either by URL or using JSON arrays")
	initCmd.Flags().StringArray("require", nil, "Package to require with a version constraint, e.g. foo/bar:1.0.0 or foo/bar=1.0.0 or \"foo/bar 1.0.0\"")
	initCmd.Flags().StringArray("require-dev", nil, "Package to require for development with a version constraint, e.g. foo/bar:1.0.0 or foo/bar=1.0.0 or \"foo/bar 1.0.0\"")
	initCmd.Flags().StringP("stability", "s", "", "Minimum stability (empty or one of: stable, RC, beta, alpha, dev)")
	initCmd.Flags().String("type", "", "Type of package (e.g. library, project, metapackage, composer-plugin)")

	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"autoload":  carapace.ActionDirectories(),
		"stability": carapace.ActionValues("RC", "alpha", "beta", "dev", "stable"),
		"type":      carapace.ActionValues("composer-plugin", "library", "metapackage", "project"),
	})
}
