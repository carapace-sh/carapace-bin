package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registered_limit_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a registered limit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registered_limit_createCmd).Standalone()

	registered_limit_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	registered_limit_createCmd.Flags().String("default-limit", "", "The default limit for the resources to assume (required)")
	registered_limit_createCmd.Flags().String("description", "", "Description of the registered limit")
	registered_limit_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	registered_limit_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	registered_limit_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	registered_limit_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	registered_limit_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	registered_limit_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	registered_limit_createCmd.Flags().String("region", "", "Region for the registered limit to affect")
	registered_limit_createCmd.Flags().String("service", "", "Service responsible for the resource to limit (required) (name or ID)")
	registered_limit_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	registered_limit_createCmd.MarkFlagRequired("default-limit")
	registered_limit_createCmd.MarkFlagRequired("service")
	registered_limitCmd.AddCommand(registered_limit_createCmd)
}
