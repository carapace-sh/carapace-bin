package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var python_upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade installed Python versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(python_upgradeCmd).Standalone()

	python_upgradeCmd.Flags().Bool("compile", false, "Compile Python's standard library to bytecode after installation")
	python_upgradeCmd.Flags().Bool("compile-bytecode", false, "Compile Python's standard library to bytecode after installation")
	python_upgradeCmd.Flags().StringP("install-dir", "i", "", "The directory Python installations are stored in")
	python_upgradeCmd.Flags().String("mirror", "", "Set the URL to use as the source for downloading Python installations")
	python_upgradeCmd.Flags().Bool("no-compile", false, "")
	python_upgradeCmd.Flags().Bool("no-compile-bytecode", false, "")
	python_upgradeCmd.Flags().String("pypy-mirror", "", "Set the URL to use as the source for downloading PyPy installations")
	python_upgradeCmd.Flags().String("python-downloads-json-url", "", "URL pointing to JSON of custom Python installations")
	python_upgradeCmd.Flags().BoolP("reinstall", "r", false, "Reinstall the latest Python patch, if it's already installed")
	python_upgradeCmd.Flag("compile").Hidden = true
	python_upgradeCmd.Flag("no-compile").Hidden = true
	python_upgradeCmd.Flag("no-compile-bytecode").Hidden = true
	pythonCmd.AddCommand(python_upgradeCmd)
	carapace.Gen(python_upgradeCmd).FlagCompletion(carapace.ActionMap{
		"install-dir": carapace.ActionDirectories(),
	})
	carapace.Gen(python_upgradeCmd).PositionalAnyCompletion(uv.ActionPythonInstallations(uv.InstallationsOpts{InstalledOnly: true}))
}
