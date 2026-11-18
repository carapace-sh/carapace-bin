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

	generateFormulaApiCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	generateFormulaApiCmd.Flags().BoolP("dry-run", "n", false, "Generate API data without writing it to files.")
	generateFormulaApiCmd.Flags().BoolP("help", "h", false, "Show this message.")
	generateFormulaApiCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	generateFormulaApiCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(generateFormulaApiCmd)
}
