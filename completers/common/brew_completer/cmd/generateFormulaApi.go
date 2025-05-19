package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generateFormulaApiCmd = &cobra.Command{
	Use:     "generate-formula-api",
	Short:   "Generate `homebrew/core` API data files for <https://formulae.brew.sh>",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateFormulaApiCmd).Standalone()

	generateFormulaApiCmd.Flags().Bool("debug", false, "Display any debugging information.")
	generateFormulaApiCmd.Flags().Bool("dry-run", false, "Generate API data without writing it to files.")
	generateFormulaApiCmd.Flags().Bool("help", false, "Show this message.")
	generateFormulaApiCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	generateFormulaApiCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(generateFormulaApiCmd)
}
