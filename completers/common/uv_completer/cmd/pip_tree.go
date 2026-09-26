package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var pip_treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Display the dependency tree for an environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pip_treeCmd).Standalone()

	pip_treeCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	pip_treeCmd.Flags().StringP("depth", "d", "255", "Maximum display depth of the dependency tree")
	pip_treeCmd.Flags().Bool("disable-pip-version-check", false, "")
	pip_treeCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	pip_treeCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	pip_treeCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	pip_treeCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	pip_treeCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	pip_treeCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	pip_treeCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	pip_treeCmd.Flags().Bool("invert", false, "Show the reverse dependencies for the given package. This flag will invert the tree and display the packages that depend on the given package")
	pip_treeCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	pip_treeCmd.Flags().Bool("no-dedupe", false, "Do not de-duplicate repeated dependencies. Usually, when a package has already displayed its dependencies, further occurrences will not re-display its dependencies, and will include a (*) to indicate it has already been shown. This flag will cause those duplicates to be repeated")
	pip_treeCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	pip_treeCmd.Flags().Bool("no-strict", false, "")
	pip_treeCmd.Flags().Bool("no-system", false, "")
	pip_treeCmd.Flags().Bool("outdated", false, "Show the latest available version of each package in the tree")
	pip_treeCmd.Flags().StringSlice("package", nil, "Display only the specified packages")
	pip_treeCmd.Flags().StringSlice("prune", nil, "Prune the given package from the display of the dependency tree")
	pip_treeCmd.Flags().StringP("python", "p", "", "The Python interpreter for which packages should be listed.")
	pip_treeCmd.Flags().Bool("reverse", false, "Show the reverse dependencies for the given package. This flag will invert the tree and display the packages that depend on the given package")
	pip_treeCmd.Flags().Bool("show-sizes", false, "Show compressed wheel sizes for packages in the tree")
	pip_treeCmd.Flags().Bool("show-version-specifiers", false, "Show the version constraint(s) imposed on each package")
	pip_treeCmd.Flags().Bool("strict", false, "Validate the Python environment, to detect packages with missing dependencies and other issues")
	pip_treeCmd.Flags().Bool("system", false, "List packages in the system Python environment")
	pip_treeCmd.Flag("disable-pip-version-check").Hidden = true
	pip_treeCmd.Flag("no-strict").Hidden = true
	pip_treeCmd.Flag("no-system").Hidden = true
	pip_treeCmd.Flag("reverse").Hidden = true
	pipCmd.AddCommand(pip_treeCmd)
	carapace.Gen(pip_treeCmd).FlagCompletion(carapace.ActionMap{
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
		"python":           uv.ActionPythonInstallations(uv.InstallationsOpts{}),
	})
}
