package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var python_findCmd = &cobra.Command{
	Use:   "find",
	Short: "Search for a Python installation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(python_findCmd).Standalone()

	python_findCmd.Flags().Bool("no-project", false, "Avoid discovering a project or workspace")
	python_findCmd.Flags().Bool("no-system", false, "")
	python_findCmd.Flags().Bool("no_workspace", false, "Avoid discovering a project or workspace")
	python_findCmd.Flags().String("python-downloads-json-url", "", "URL pointing to JSON of custom Python installations")
	python_findCmd.Flags().Bool("resolve-links", false, "Resolve symlinks in the output path")
	python_findCmd.Flags().String("script", "", "Find the environment for a Python script, rather than the current project")
	python_findCmd.Flags().Bool("show-version", false, "Show the Python version that would be used instead of the path to the interpreter")
	python_findCmd.Flags().Bool("system", false, "Only find system Python interpreters")
	python_findCmd.Flag("no-system").Hidden = true
	python_findCmd.Flag("no_workspace").Hidden = true
	pythonCmd.AddCommand(python_findCmd)
	carapace.Gen(python_findCmd).FlagCompletion(carapace.ActionMap{
		"script": carapace.ActionFiles(),
	})
	carapace.Gen(python_findCmd).PositionalCompletion(uv.ActionPythonInstallations(uv.InstallationsOpts{InstalledOnly: true}))
}
