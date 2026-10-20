// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var testToolCmd = &cobra.Command{
	Use:   "test-tool",
	Short: "Test that a tool installs and runs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testToolCmd).Standalone()

	testToolCmd.Flags().BoolP("all", "a", false, "Test every registry tool that defines a `test` and supports this OS")
	testToolCmd.Flags().Bool("all-config", false, "Test all tools specified in config files")
	testToolCmd.Flags().BoolP("help", "h", false, "Print help")
	testToolCmd.Flags().Bool("include-non-defined", false, "Also test registry tools that have no `test` entry by running `<tool> --version`")
	testToolCmd.Flags().StringP("jobs", "j", "", "Number of tools to test in parallel (default 4)")
	testToolCmd.Flags().Bool("raw", false, "Connect the install commands' stdin, stdout, and stderr to the terminal")
	rootCmd.AddCommand(testToolCmd)
}
