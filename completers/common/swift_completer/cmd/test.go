package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/swift_completer/cmd/common"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Build and run tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()
	testCmd.Flags().SetInterspersed(false)

	common.AddPackageFlags(testCmd)

	testCmd.Flags().String("attachments-path", "", "Path where attachments should be written")
	testCmd.Flags().Bool("disable-code-coverage", false, "Disable code coverage")
	testCmd.Flags().Bool("disable-swift-testing", false, "Disable support for Swift Testing")
	testCmd.Flags().Bool("disable-testable-imports", false, "Disable testable imports")
	testCmd.Flags().Bool("disable-xctest", false, "Disable support for XCTest")
	testCmd.Flags().Bool("enable-code-coverage", false, "Enable code coverage")
	testCmd.Flags().Bool("enable-swift-testing", false, "Enable support for Swift Testing")
	testCmd.Flags().Bool("enable-testable-imports", false, "Enable testable imports")
	testCmd.Flags().Bool("enable-xctest", false, "Enable support for XCTest")
	testCmd.Flags().String("filter", "", "Run test cases that match a regular expression")
	testCmd.Flags().BoolP("help", "h", false, "Show help information")
	testCmd.Flags().BoolP("list-tests", "l", false, "Lists test methods in specifier format")
	testCmd.Flags().Bool("no-parallel", false, "Do not run tests in parallel")
	testCmd.Flags().Int("num-workers", 0, "Number of tests to execute in parallel")
	testCmd.Flags().Bool("parallel", false, "Run the tests in parallel")
	testCmd.Flags().Bool("show-code-coverage-path", false, "Print the path of the exported code coverage JSON file")
	testCmd.Flags().Bool("show-codecov-path", false, "Print the path of the exported code coverage JSON file")
	testCmd.Flags().Bool("show-coverage-path", false, "Print the path of the exported code coverage JSON file")
	testCmd.Flags().String("skip", "", "Skip test cases that match a regular expression")
	testCmd.Flags().Bool("skip-build", false, "Skip building the test target")
	testCmd.Flags().StringP("specifier", "s", "", "Specify a test specifier")
	testCmd.Flags().Bool("version", false, "Show the version")
	testCmd.Flags().String("xunit-output", "", "Path where the xUnit xml file should be generated")

	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"attachments-path": carapace.ActionDirectories(),
		"xunit-output":     carapace.ActionFiles(),
	})

}
