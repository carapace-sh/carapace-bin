package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nilaway [-flag] [package]",
	Short: "Static Analysis tool to detect potential Nil panics in Go code",
	Long:  "https://github.com/uber-go/nilaway",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().BoolS("V", "V", false, "print version and exit")
	rootCmd.Flags().StringS("c", "c", "", "display offending line with this many lines of context")
	rootCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "write CPU profile to this file")
	rootCmd.Flags().StringS("debug", "debug", "", "debug flags, any subset of \"fpstv\"")
	rootCmd.Flags().StringS("exclude-errors-in-files", "exclude-errors-in-files", "", "A comma-separated list of file prefixes to exclude from error reporting")
	rootCmd.Flags().StringS("exclude-file-docstrings", "exclude-file-docstrings", "", "Comma-separated list of docstrings to exclude from analysis")
	rootCmd.Flags().StringS("exclude-pkgs", "exclude-pkgs", "", "Comma-separated list of packages to exclude from analysis")
	rootCmd.Flags().BoolS("fix", "fix", false, "apply all suggested fixes")
	rootCmd.Flags().BoolS("flags", "flags", false, "print analyzer flags in JSON")
	rootCmd.Flags().StringS("include-errors-in-files", "include-errors-in-files", "", "A comma-separated list of file prefixes to report errors")
	rootCmd.Flags().StringS("include-pkgs", "include-pkgs", "", "Comma-separated list of packages to analyze")
	rootCmd.Flags().BoolS("json", "json", false, "emit JSON output")
	rootCmd.Flags().StringS("memprofile", "memprofile", "", "write memory profile to this file")
	rootCmd.Flags().StringS("pretty-print", "pretty-print", "", "Pretty print the error messages")
	rootCmd.Flags().BoolS("test", "test", false, "indicates whether test files should be analyzed, too")
	rootCmd.Flags().StringS("trace", "trace", "", "write trace log to this file")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile":              carapace.ActionFiles(),
		"debug":                   carapace.ActionValues(), // TODO
		"exclude-errors-in-files": carapace.ActionFiles().UniqueList(","),
		"exclude-pkgs":            golang.ActionPackages().UniqueList(","),
		"include-errors-in-files": carapace.ActionFiles().UniqueList(","),
		"include-pkgs":            golang.ActionPackages().UniqueList(","),
		"memprofile":              carapace.ActionFiles(),
		"pretty-print":            carapace.ActionValues(), // TODO
		"trace":                   carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		golang.ActionPackages(),
	)
}
