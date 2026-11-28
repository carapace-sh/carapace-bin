package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clangCheckCmd = &cobra.Command{
	Use:   "clang-check",
	Short: "Run clang-tidy check under current folder",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clangCheckCmd).Standalone()

	clangCheckCmd.Flags().Bool("all-files", false, "Run clang-tidy with all files")
	clangCheckCmd.Flags().StringArray("exclude-paths", nil, "Exclude extra files besides of the project dir")
	clangCheckCmd.Flags().Bool("exit-code", false, "Exit with code based on the results of the code analysis")
	clangCheckCmd.Flags().StringArray("include-paths", nil, "Include extra files besides of the project dir")
	clangCheckCmd.Flags().String("run-clang-tidy-options", "", "Optional arguments passed to run-clang-tidy.py")
	clangCheckCmd.Flags().String("run-clang-tidy-py", "", "Path to run-clang-tidy.py")

	rootCmd.AddCommand(clangCheckCmd)

	carapace.Gen(clangCheckCmd).FlagCompletion(carapace.ActionMap{
		"exclude-paths":     carapace.ActionFiles(),
		"include-paths":     carapace.ActionFiles(),
		"run-clang-tidy-py": carapace.ActionFiles(),
	})
}
