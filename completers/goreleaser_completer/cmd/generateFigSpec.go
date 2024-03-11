package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generateFigSpecCmd = &cobra.Command{
	Use:     "generate-fig-spec",
	Short:   "Generate a fig spec",
	Aliases: []string{"generateFigSpec", "genFigSpec"},
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateFigSpecCmd).Standalone()

	generateFigSpecCmd.Flags().Bool("include-hidden", false, "Include hidden commands in generated Fig autocomplete spec")
	rootCmd.AddCommand(generateFigSpecCmd)
}
