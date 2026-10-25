package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run",
	Short:   "Run KCL codes",
	Aliases: []string{"r"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().StringArrayP("argument", "D", []string{}, "Specify the top-level argument")
	runCmd.Flags().StringP("branch", "b", "", "Specify the branch for the Git artifact")
	runCmd.Flags().StringP("commit", "c", "", "Specify the commit for the Git artifact")
	runCmd.Flags().BoolP("debug", "d", false, "Run in debug mode")
	runCmd.Flags().BoolP("disable_none", "n", false, "Disable dumping None values")
	runCmd.Flags().String("error_format", "", "Specify the diagnostic output format")
	runCmd.Flags().StringArrayP("external", "E", []string{}, "Specify the mapping of package name and path where the package is located")
	runCmd.Flags().String("format", "yaml", "Specify the output format")
	runCmd.Flags().String("git", "", "Specify the KCL module git url")
	runCmd.Flags().Bool("no_style", false, "Set to prohibit output of command line waiting styles, including colors, etc.")
	runCmd.Flags().String("oci", "", "Specify the KCL module oci url")
	runCmd.Flags().StringP("output", "o", "", "Specify the YAML/JSON output file path")
	runCmd.Flags().StringArrayP("overrides", "O", []string{}, "Specify the configuration override path and value")
	runCmd.Flags().StringArrayP("path_selector", "S", []string{}, "Specify the path selectors")
	runCmd.Flags().BoolP("quiet", "q", false, "Set the quiet mode (no output)")
	runCmd.Flags().StringArrayP("setting", "Y", []string{}, "Specify the command line setting files")
	runCmd.Flags().BoolP("show_hidden", "H", false, "Display hidden attributes")
	runCmd.Flags().BoolP("sort_keys", "k", false, "Sort output result keys")
	runCmd.Flags().BoolP("strict_range_check", "r", false, "Do perform strict numeric range checks")
	runCmd.Flags().StringP("tag", "t", "", "Specify the tag for the OCI or Git artifact")
	runCmd.Flags().BoolP("vendor", "V", false, "Run in vendor mode")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"error_format": carapace.ActionValues("arcanist", "pretty", "sarif", "short"),
		"format":       carapace.ActionValues("json", "toml", "xml", "yaml"),
		"output":       carapace.ActionFiles(),
		"setting":      carapace.ActionFiles(".yaml", ".yml"),
	})

	carapace.Gen(runCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".k"),
	)
}
