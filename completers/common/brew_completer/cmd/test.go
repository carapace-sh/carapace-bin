package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "test",
	Short:   "Run the test method provided by an installed formula",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().Bool("HEAD", false, "Test the HEAD version of a formula.")
	testCmd.Flags().Bool("debug", false, "Display any debugging information.")
	testCmd.Flags().Bool("force", false, "Test formulae even if they are unlinked.")
	testCmd.Flags().Bool("help", false, "Show this message.")
	testCmd.Flags().Bool("keep-tmp", false, "Retain the temporary files created for the test.")
	testCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	testCmd.Flags().Bool("retry", false, "Retry if a testing fails.")
	testCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(testCmd)
}
