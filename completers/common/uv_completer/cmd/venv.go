package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var venvCmd = &cobra.Command{
	Use:     "venv",
	Short:   "Create a virtual environment",
	Aliases: []string{"virtualenv", "v"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(venvCmd).Standalone()

	venvCmd.Flags().Bool("allow-existing", false, "Preserve any existing files or directories at the target path")
	venvCmd.Flags().BoolP("clear", "c", false, "Remove any existing files or directories at the target path [env: UV_VENV_CLEAR=]")
	venvCmd.Flags().String("default-index", "", "The default package index (by default: <https://pypi.org/simple>)")
	venvCmd.Flags().String("exclude-newer", "", "Limit candidate packages to those that were uploaded prior to the given date")
	venvCmd.Flags().StringSlice("exclude-newer-package", nil, "Limit candidate packages for specific packages to those that were uploaded prior to the given date")
	venvCmd.Flags().StringSlice("extra-index-url", nil, "(Deprecated: use `--index` instead) Extra URLs of package indexes to use, in addition to `--index-url`")
	venvCmd.Flags().StringSliceP("find-links", "f", nil, "Locations to search for candidate distributions, in addition to those found in the registry indexes")
	venvCmd.Flags().Bool("force", false, "Allow `--clear` to remove a non-virtual environment directory")
	venvCmd.Flags().StringSlice("index", nil, "The indexes to use when resolving dependencies, in addition to the default index")
	venvCmd.Flags().String("index-strategy", "", "The strategy to use when resolving against multiple index URLs")
	venvCmd.Flags().StringP("index-url", "i", "", "(Deprecated: use `--default-index` instead) The URL of the Python package index (by default: <https://pypi.org/simple>)")
	venvCmd.Flags().String("keyring-provider", "", "Attempt to use `keyring` for authentication for index URLs")
	venvCmd.Flags().String("link-mode", "", "The method to use when installing packages from the global cache")
	venvCmd.Flags().Bool("no-clear", false, "Fail without prompting if any existing files or directories are present at the target path")
	venvCmd.Flags().Bool("no-index", false, "Ignore the registry index (e.g., PyPI), instead relying on direct URL dependencies and those provided via `--find-links`")
	venvCmd.Flags().Bool("no-pip", false, "")
	venvCmd.Flags().Bool("no-project", false, "Avoid discovering a project or workspace")
	venvCmd.Flags().Bool("no-refresh", false, "")
	venvCmd.Flags().Bool("no-relocatable", false, "Don't make the virtual environment relocatable")
	venvCmd.Flags().Bool("no-seed", false, "")
	venvCmd.Flags().Bool("no-setuptools", false, "")
	venvCmd.Flags().Bool("no-system", false, "This flag is included for compatibility only, it has no effect")
	venvCmd.Flags().Bool("no-wheel", false, "")
	venvCmd.Flags().Bool("no-workspace", false, "Avoid discovering a project or workspace")
	venvCmd.Flags().String("prompt", "", "Provide an alternative prompt prefix for the virtual environment.")
	venvCmd.Flags().StringP("python", "p", "", "The Python interpreter to use for the virtual environment.")
	venvCmd.Flags().Bool("refresh", false, "Refresh all cached data")
	venvCmd.Flags().StringSlice("refresh-package", nil, "Refresh cached data for a specific package")
	venvCmd.Flags().Bool("relocatable", false, "Make the virtual environment relocatable [env: UV_VENV_RELOCATABLE=]")
	venvCmd.Flags().Bool("seed", false, "Install seed packages (one or more of: `pip`, `setuptools`, and `wheel`) into the virtual environment [env: UV_VENV_SEED=]")
	venvCmd.Flags().Bool("system", false, "Ignore virtual environments when searching for the Python interpreter")
	venvCmd.Flags().Bool("system-site-packages", false, "Give the virtual environment access to the system site packages directory")
	venvCmd.Flag("no-clear").Hidden = true
	venvCmd.Flag("no-pip").Hidden = true
	venvCmd.Flag("no-refresh").Hidden = true
	venvCmd.Flag("no-relocatable").Hidden = true
	venvCmd.Flag("no-seed").Hidden = true
	venvCmd.Flag("no-setuptools").Hidden = true
	venvCmd.Flag("no-system").Hidden = true
	venvCmd.Flag("no-wheel").Hidden = true
	venvCmd.Flag("no-workspace").Hidden = true
	venvCmd.Flag("system").Hidden = true
	rootCmd.AddCommand(venvCmd)
	carapace.Gen(venvCmd).FlagCompletion(carapace.ActionMap{
		"index-strategy":   uv.ActionIndexStrategies(),
		"keyring-provider": uv.ActionKeyringProviders(),
		"link-mode":        uv.ActionLinkModes(),
		"python":           uv.ActionPythonInstallations(uv.InstallationsOpts{}),
	})
	carapace.Gen(venvCmd).PositionalCompletion(carapace.ActionDirectories())
}
