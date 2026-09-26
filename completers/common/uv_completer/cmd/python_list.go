package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var python_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List the available Python installations",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(python_listCmd).Standalone()

	python_listCmd.Flags().Bool("all-arches", false, "List Python downloads for all architectures")
	python_listCmd.Flags().Bool("all-platforms", false, "List Python downloads for all platforms")
	python_listCmd.Flags().Bool("all-versions", false, "List all Python versions, including old patch versions")
	python_listCmd.Flags().Bool("all_architectures", false, "List Python downloads for all architectures")
	python_listCmd.Flags().Bool("only-downloads", false, "Only show available Python downloads")
	python_listCmd.Flags().Bool("only-installed", false, "Only show installed Python versions")
	python_listCmd.Flags().String("output-format", "text", "Select the output format")
	python_listCmd.Flags().String("python-downloads-json-url", "", "URL pointing to JSON of custom Python installations")
	python_listCmd.Flags().Bool("show-urls", false, "Show the URLs of available Python downloads")
	python_listCmd.Flag("all_architectures").Hidden = true
	pythonCmd.AddCommand(python_listCmd)
	carapace.Gen(python_listCmd).FlagCompletion(carapace.ActionMap{
		"output-format": carapace.ActionValuesDescribed(
			"text", "Plain text (for humans)",
			"json", "JSON (for computers)",
		),
	})
	carapace.Gen(python_listCmd).PositionalCompletion(uv.ActionPythonInstallations(uv.InstallationsOpts{}))
}
