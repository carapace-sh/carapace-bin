package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var test_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists test methods in specifier format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(test_listCmd).Standalone()
	test_listCmd.Flags().SetInterspersed(false)

	test_listCmd.Flags().String("attachments-path", "", "Path where attachments should be written")
	test_listCmd.Flags().Bool("disable-swift-testing", false, "Disable support for Swift Testing")
	test_listCmd.Flags().Bool("disable-xctest", false, "Disable support for XCTest")
	test_listCmd.Flags().Bool("enable-swift-testing", false, "Enable support for Swift Testing")
	test_listCmd.Flags().Bool("enable-xctest", false, "Enable support for XCTest")
	test_listCmd.Flags().BoolP("help", "h", false, "Show help information")
	test_listCmd.Flags().Bool("skip-build", false, "Skip building the test target")
	test_listCmd.Flags().Bool("version", false, "Show the version")

	testCmd.AddCommand(test_listCmd)

	carapace.Gen(test_listCmd).FlagCompletion(carapace.ActionMap{
		"attachments-path": carapace.ActionDirectories(),
	})
}
