package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var test_lastCmd = &cobra.Command{
	Use:   "last",
	Short: "Show the last test run",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(test_lastCmd).Standalone()
	test_lastCmd.Flags().SetInterspersed(false)

	test_lastCmd.Flags().BoolP("help", "h", false, "Show help information")
	test_lastCmd.Flags().Bool("version", false, "Show the version")

	testCmd.AddCommand(test_lastCmd)
}
