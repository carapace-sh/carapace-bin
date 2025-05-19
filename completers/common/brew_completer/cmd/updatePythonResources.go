package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updatePythonResourcesCmd = &cobra.Command{
	Use:     "update-python-resources",
	Short:   "Update versions for PyPI resource blocks in <formula>",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updatePythonResourcesCmd).Standalone()

	updatePythonResourcesCmd.Flags().Bool("debug", false, "Display any debugging information.")
	updatePythonResourcesCmd.Flags().Bool("exclude-packages", false, "Exclude these packages when finding resources.")
	updatePythonResourcesCmd.Flags().Bool("extra-packages", false, "Include these additional packages when finding resources.")
	updatePythonResourcesCmd.Flags().Bool("help", false, "Show this message.")
	updatePythonResourcesCmd.Flags().Bool("ignore-non-pypi-packages", false, "Don't fail if <formula> is not a PyPI package.")
	updatePythonResourcesCmd.Flags().Bool("package-name", false, "Use the specified <package-name> when finding resources for <formula>. If no package name is specified, it will be inferred from the formula's stable URL.")
	updatePythonResourcesCmd.Flags().Bool("print-only", false, "Print the updated resource blocks instead of changing <formula>.")
	updatePythonResourcesCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	updatePythonResourcesCmd.Flags().Bool("silent", false, "Suppress any output.")
	updatePythonResourcesCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	updatePythonResourcesCmd.Flags().Bool("version", false, "Use the specified <version> when finding resources for <formula>. If no version is specified, the current version for <formula> will be used.")
	rootCmd.AddCommand(updatePythonResourcesCmd)
}
