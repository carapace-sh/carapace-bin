package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Tests that a book's Rust code samples compile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().StringP("chapter", "c", "", "")
	testCmd.Flags().BoolP("help", "h", false, "Print help")
	testCmd.Flags().StringSliceP("library-path", "L", nil, "A comma-separated list of directories to add to the crate search path when building tests")
	testCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"library-path": carapace.ActionDirectories().List(","),
	})

	carapace.Gen(testCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
