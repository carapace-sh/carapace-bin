package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var formulaCmd = &cobra.Command{
	Use:     "formula",
	Short:   "Display the path where <formula> is located",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formulaCmd).Standalone()

	formulaCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	formulaCmd.Flags().BoolP("help", "h", false, "Show this message.")
	formulaCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	formulaCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(formulaCmd)

	carapace.Gen(formulaCmd).PositionalAnyCompletion(
		brew.ActionAllFormulae(),
	)
}
