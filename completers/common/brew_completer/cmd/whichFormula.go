package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var whichFormulaCmd = &cobra.Command{
	Use:     "which-formula",
	Short:   "Show which formula(e) provides the given command",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whichFormulaCmd).Standalone()

	whichFormulaCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	whichFormulaCmd.Flags().BoolP("help", "h", false, "Show this message.")
	whichFormulaCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	whichFormulaCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(whichFormulaCmd)

	carapace.Gen(whichFormulaCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
