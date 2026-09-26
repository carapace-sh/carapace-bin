package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var python_pinCmd = &cobra.Command{
	Use:   "pin",
	Short: "Pin to a specific Python version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(python_pinCmd).Standalone()

	python_pinCmd.Flags().Bool("global", false, "Update the global Python version pin")
	python_pinCmd.Flags().Bool("no-project", false, "Avoid validating the Python pin is compatible with the project or workspace")
	python_pinCmd.Flags().Bool("no-resolved", false, "")
	python_pinCmd.Flags().Bool("no-workspace", false, "Avoid validating the Python pin is compatible with the project or workspace")
	python_pinCmd.Flags().String("python-downloads-json-url", "", "URL pointing to JSON of custom Python installations")
	python_pinCmd.Flags().Bool("resolved", false, "Write the resolved Python interpreter path instead of the request")
	python_pinCmd.Flags().Bool("rm", false, "Remove the Python version pin")
	python_pinCmd.Flag("no-resolved").Hidden = true
	python_pinCmd.Flag("no-workspace").Hidden = true
	pythonCmd.AddCommand(python_pinCmd)
	carapace.Gen(python_pinCmd).PositionalCompletion(uv.ActionPythonInstallations(uv.InstallationsOpts{InstalledOnly: true}))
}
