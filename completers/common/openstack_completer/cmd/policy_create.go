package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var policy_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new policy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(policy_createCmd).Standalone()

	policy_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	policy_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	policy_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	policy_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	policy_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	policy_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	policy_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	policy_createCmd.Flags().String("type", "", "New MIME type of the policy rules file (defaults to application/json)")
	policy_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	policyCmd.AddCommand(policy_createCmd)
}
