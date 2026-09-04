package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var policy_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display policy details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policy_showCmd).Standalone()

	policy_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	policy_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	policy_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	policy_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	policy_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	policy_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	policy_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	policy_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	policyCmd.AddCommand(policy_showCmd)
}
