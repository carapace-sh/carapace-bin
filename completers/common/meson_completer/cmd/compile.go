package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Build the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compileCmd).Standalone()

	compileCmd.Flags().StringS("C", "C", "", "directory to cd into before running")
	compileCmd.Flags().Bool("clean", false, "Clean the build directory.")
	compileCmd.Flags().StringP("jobs", "j", "", "The number of worker jobs to run (if supported).")
	compileCmd.Flags().StringP("load-average", "l", "", "The system load average to try to maintain (if supported).")
	compileCmd.Flags().String("ninja-args", "", "Arguments to pass to `ninja` (applied only on `ninja` backend).")
	compileCmd.Flags().BoolP("verbose", "v", false, "Show more verbose output.")
	compileCmd.Flags().String("vs-args", "", "Arguments to pass to `msbuild` (applied only on `vs` backend).")
	compileCmd.Flags().String("xcode-args", "", "Arguments to pass to `xcodebuild` (applied only on `xcode` backend).")
	applyCommonFlagCompletion(compileCmd)
	rootCmd.AddCommand(compileCmd)
}
