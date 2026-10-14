package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_addDependencyCmd = &cobra.Command{
	Use:   "add-dependency",
	Short: "Add a package dependency to the manifest",
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_addDependencyCmd).Standalone()
	package_addDependencyCmd.Flags().SetInterspersed(false)

	package_addDependencyCmd.Flags().String("branch", "", "The branch of the package to depend on")
	package_addDependencyCmd.Flags().String("exact", "", "The exact package version to depend on")
	package_addDependencyCmd.Flags().String("from", "", "The package version to depend on (up to the next major version)")
	package_addDependencyCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_addDependencyCmd.Flags().String("revision", "", "The specific package revision to depend on")
	package_addDependencyCmd.Flags().String("to", "", "Specify upper bound on the package version range (exclusive)")
	package_addDependencyCmd.Flags().String("type", "", "Specify dependency type")
	package_addDependencyCmd.Flags().String("up-to-next-minor-from", "", "The package version to depend on (up to the next minor version)")
	package_addDependencyCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_addDependencyCmd)

	carapace.Gen(package_addDependencyCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("url", "path", "registry"),
	})
}
