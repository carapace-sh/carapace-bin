package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runTestsCmd = &cobra.Command{
	Use:   "run-tests",
	Short: "Run unit tests of the project using mir interpreter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runTestsCmd).Standalone()

	rootCmd.AddCommand(runTestsCmd)

	carapace.Gen(runTestsCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
