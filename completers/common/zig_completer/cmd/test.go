package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Perform unit testing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().StringP("test-filter", "", "", "Skip tests that do not match any filter")
	testCmd.Flags().StringP("test-name-prefix", "", "", "Prefix all tests with given string")
	testCmd.Flags().Bool("test-cmd-bin", false, "Appends test binary path to test-cmd argv")
	testCmd.Flags().Bool("test-evented-io", false, "Runs the test in evented I/O mode")
	testCmd.Flags().Bool("test-no-exec", false, "Compiles test binary without running it")
	testCmd.Flags().StringP("name", "", "", "Override root name")
	testCmd.Flags().Bool("strip", false, "Omit debug symbols")
	testCmd.Flags().StringP("target", "", "", "<arch><sub>-<os>-<abi>")
	testCmd.Flags().StringP("mcpu", "", "", "Specify target CPU and feature set")
	testCmd.Flags().StringP("O", "", "", "Choose what to optimize for")
	testCmd.Flags().BoolP("help", "h", false, "Print help")

	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"O": carapace.ActionValues(
			"Debug",
			"ReleaseSafe",
			"ReleaseFast",
			"ReleaseSmall",
		),
	})

	carapace.Gen(testCmd).PositionalCompletion(
		carapace.ActionFiles(".zig"),
	)
}
