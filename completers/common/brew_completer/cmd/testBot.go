package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var testBotCmd = &cobra.Command{
	Use:     "test-bot",
	Short:   "Test the full lifecycle of a Homebrew change to a tap (Git repository)",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testBotCmd).Standalone()

	testBotCmd.Flags().Bool("added-formulae", false, "Only test added formulae.")
	testBotCmd.Flags().Bool("debug", false, "Display any debugging information.")
	testBotCmd.Flags().Bool("dry-run", false, "Print what would be done rather than doing it.")
	testBotCmd.Flags().Bool("eval-all", false, "Evaluate all formulae and casks in the tap.")
	testBotCmd.Flags().Bool("fail-fast", false, "Exit early on the first failing test.")
	testBotCmd.Flags().Bool("help", false, "Show this message.")
	testBotCmd.Flags().Bool("keep-old", false, "Keep previous versions of formulae built by `brew test-bot`.")
	testBotCmd.Flags().Bool("no-dependents", false, "Don't test dependents.")
	testBotCmd.Flags().Bool("only-cleanup-after", false, "Only run cleanup after testing.")
	testBotCmd.Flags().Bool("only-cleanup-before", false, "Only run cleanup before testing.")
	testBotCmd.Flags().Bool("only-formulae", false, "Only test formulae.")
	testBotCmd.Flags().Bool("only-setup", false, "Only run setup steps.")
	testBotCmd.Flags().Bool("only-tap-syntax", false, "Only test tap syntax.")
	testBotCmd.Flags().Bool("publish", false, "Publish built bottles.")
	testBotCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	testBotCmd.Flags().Bool("skip-dependents", false, "Skip testing dependents.")
	testBotCmd.Flags().Bool("tap", false, "Target tap repository (default: `homebrew/core`).")
	testBotCmd.Flags().Bool("testing-formulae", false, "Only test testing formulae.")
	testBotCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(testBotCmd)
}
