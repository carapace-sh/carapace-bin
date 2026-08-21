package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var implied_role_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates an association between prior and implied roles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(implied_role_createCmd).Standalone()

	implied_role_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	implied_role_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	implied_role_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	implied_role_createCmd.Flags().String("implied-role", "", "<role> (name or ID) implied by another role")
	implied_role_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	implied_role_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	implied_role_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	implied_role_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	implied_role_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	implied_role_createCmd.MarkFlagRequired("implied-role")
	implied_roleCmd.AddCommand(implied_role_createCmd)
}
