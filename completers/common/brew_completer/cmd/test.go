package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
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
	testCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	testCmd.Flags().BoolP("force", "f", false, "Test formulae even if they are unlinked.")
	testCmd.Flags().BoolP("help", "h", false, "Show this message.")
	testCmd.Flags().Bool("keep-tmp", false, "Retain the temporary files created for the test.")
	testCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	testCmd.Flags().Bool("retry", false, "Retry if a testing fails.")
	testCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).PositionalAnyCompletion(
		action.ActionList(testCmd).FilterArgs(),
	)
}
