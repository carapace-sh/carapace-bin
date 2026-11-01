package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "KCL test tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().Bool("fail-fast", false, "Stop the KCL test run on the first failed test case")
	testCmd.Flags().String("run", "", "Run only tests matching the regular expression")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
