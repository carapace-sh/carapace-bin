package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var styleCmd = &cobra.Command{
	Use:     "style",
	Short:   "Check formulae or files for conformance to Homebrew style guidelines",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(styleCmd).Standalone()

	styleCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	styleCmd.Flags().Bool("debug", false, "Display any debugging information.")
	styleCmd.Flags().Bool("except-cops", false, "Specify a comma-separated <cops> list to skip checking for violations of the listed RuboCop cops.")
	styleCmd.Flags().Bool("fix", false, "Fix style violations automatically using RuboCop's auto-correct feature.")
	styleCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	styleCmd.Flags().Bool("help", false, "Show this message.")
	styleCmd.Flags().Bool("only-cops", false, "Specify a comma-separated <cops> list to check for violations of only the listed RuboCop cops.")
	styleCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	styleCmd.Flags().Bool("reset-cache", false, "Reset the RuboCop cache.")
	styleCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(styleCmd)
}
