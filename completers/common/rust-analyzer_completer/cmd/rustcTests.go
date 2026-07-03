package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rustcTestsCmd = &cobra.Command{
	Use:   "rustc-tests",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rustcTestsCmd).Standalone()

	rustcTestsCmd.Flags().String("filter", "", "Only run tests with filter as substring")
	rootCmd.AddCommand(rustcTestsCmd)

	carapace.Gen(rustcTestsCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
