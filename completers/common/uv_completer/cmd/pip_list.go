package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var pip_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List, in tabular format, packages installed in an environment",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pip_listCmd).Standalone()

	pip_listCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	pip_listCmd.Flags().Bool("disable-pip-version-check", false, "")
	pip_listCmd.Flags().BoolP("editable", "e", false, "Only include editable projects")
	pip_listCmd.Flags().StringSlice("exclude", nil, "Exclude the specified package(s) from the output")
	pip_listCmd.Flags().Bool("exclude-editable", false, "Exclude any editable packages from output")
	pip_listCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	pip_listCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	pip_listCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	pip_listCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	pip_listCmd.Flags().String("format", "columns", "Select the output format")
	pip_listCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	pip_listCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	pip_listCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	pip_listCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	pip_listCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	pip_listCmd.Flags().Bool("no-outdated", false, "")
	pip_listCmd.Flags().Bool("no-strict", false, "")
	pip_listCmd.Flags().Bool("no-system", false, "")
	pip_listCmd.Flags().Bool("outdated", false, "List outdated packages")
	pip_listCmd.Flags().String("prefix", "", "List packages from the specified `--prefix` directory")
	pip_listCmd.Flags().StringP("python", "p", "", "The Python interpreter for which packages should be listed.")
	pip_listCmd.Flags().Bool("strict", false, "Validate the Python environment, to detect packages with missing dependencies and other issues")
	pip_listCmd.Flags().Bool("system", false, "List packages in the system Python environment")
	pip_listCmd.Flags().StringP("target", "t", "", "List packages from the specified `--target` directory")
	pip_listCmd.Flag("disable-pip-version-check").Hidden = true
	pip_listCmd.Flag("no-outdated").Hidden = true
	pip_listCmd.Flag("no-strict").Hidden = true
	pip_listCmd.Flag("no-system").Hidden = true
	pipCmd.AddCommand(pip_listCmd)
	carapace.Gen(pip_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValuesDescribed(
			"columns", "Display the list of packages in a human-readable table",
			"freeze", "Display the list of packages in a `pip freeze`-like format, with one package per line alongside its version",
			"json", "Display the list of packages in a machine-readable JSON format",
		),
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
		"prefix":           carapace.ActionDirectories(),
		"python":           uv.ActionPythonInstallations(uv.InstallationsOpts{}),
		"target":           carapace.ActionDirectories(),
	})
}
