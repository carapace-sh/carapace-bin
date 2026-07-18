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

	whichFormulaCmd.Flags().Bool("debug", false, "Display any debugging information.")
	whichFormulaCmd.Flags().Bool("help", false, "Show this message.")
	whichFormulaCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	whichFormulaCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(whichFormulaCmd)
}
