package cmd

import (
	"github.com/carapace-sh/carapace"
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

	formulaCmd.Flags().Bool("debug", false, "Display any debugging information.")
	formulaCmd.Flags().Bool("help", false, "Show this message.")
	formulaCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	formulaCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(formulaCmd)
}
