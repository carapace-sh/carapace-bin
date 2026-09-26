package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var python_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall Python versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(python_uninstallCmd).Standalone()

	python_uninstallCmd.Flags().Bool("all", false, "Uninstall all managed Python versions")
	python_uninstallCmd.Flags().StringP("install-dir", "i", "", "The directory where the Python was installed")
	pythonCmd.AddCommand(python_uninstallCmd)
	carapace.Gen(python_uninstallCmd).FlagCompletion(carapace.ActionMap{
		"install-dir": carapace.ActionDirectories(),
	})
	carapace.Gen(python_uninstallCmd).PositionalAnyCompletion(uv.ActionPythonInstallations(uv.InstallationsOpts{InstalledOnly: true}))
}
