package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Lint KCL codes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lintCmd).Standalone()

	lintCmd.Flags().StringArrayP("argument", "D", []string{}, "Specify the top-level argument")
	lintCmd.Flags().StringP("branch", "b", "", "Specify the branch for the Git artifact")
	lintCmd.Flags().StringP("commit", "c", "", "Specify the commit for the Git artifact")
	lintCmd.Flags().BoolP("debug", "d", false, "Run in debug mode")
	lintCmd.Flags().BoolP("disable_none", "n", false, "Disable dumping None values")
	lintCmd.Flags().String("error_format", "", "Specify the diagnostic output format")
	lintCmd.Flags().StringArrayP("external", "E", []string{}, "Specify the mapping of package name and path where the package is located")
	lintCmd.Flags().String("format", "yaml", "Specify the output format")
	lintCmd.Flags().String("git", "", "Specify the KCL module git url")
	lintCmd.Flags().Bool("no_style", false, "Set to prohibit output of command line waiting styles, including colors, etc.")
	lintCmd.Flags().String("oci", "", "Specify the KCL module oci url")
	lintCmd.Flags().StringP("output", "o", "", "Specify the YAML/JSON output file path")
	lintCmd.Flags().StringArrayP("overrides", "O", []string{}, "Specify the configuration override path and value")
	lintCmd.Flags().StringArrayP("path_selector", "S", []string{}, "Specify the path selectors")
	lintCmd.Flags().BoolP("quiet", "q", false, "Set the quiet mode (no output)")
	lintCmd.Flags().StringArrayP("setting", "Y", []string{}, "Specify the command line setting files")
	lintCmd.Flags().BoolP("show_hidden", "H", false, "Display hidden attributes")
	lintCmd.Flags().BoolP("sort_keys", "k", false, "Sort output result keys")
	lintCmd.Flags().BoolP("strict_range_check", "r", false, "Do perform strict numeric range checks")
	lintCmd.Flags().StringP("tag", "t", "", "Specify the tag for the OCI or Git artifact")
	lintCmd.Flags().BoolP("vendor", "V", false, "Run in vendor mode")
	rootCmd.AddCommand(lintCmd)

	carapace.Gen(lintCmd).FlagCompletion(carapace.ActionMap{
		"error_format": carapace.ActionValues("arcanist", "pretty", "sarif", "short"),
		"format":       carapace.ActionValues("json", "toml", "xml", "yaml"),
		"output":       carapace.ActionFiles(),
		"setting":      carapace.ActionFiles(".yaml", ".yml"),
	})

	carapace.Gen(lintCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".k"),
	)
}
