package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var registered_limit_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update information about a registered limit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(registered_limit_setCmd).Standalone()

	registered_limit_setCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	registered_limit_setCmd.Flags().String("default-limit", "", "The default limit for the resources to assume")
	registered_limit_setCmd.Flags().String("description", "", "Description to update of the registered limit")
	registered_limit_setCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	registered_limit_setCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	registered_limit_setCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	registered_limit_setCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	registered_limit_setCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	registered_limit_setCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	registered_limit_setCmd.Flags().String("region", "", "Region for the registered limit to affect.")
	registered_limit_setCmd.Flags().String("resource-name", "", "Resource to be updated responsible for the resource to limit.")
	registered_limit_setCmd.Flags().String("service", "", "Service to be updated responsible for the resource to limit (name or ID).")
	registered_limit_setCmd.Flags().String("variable", "", "==SUPPRESS==")
	registered_limitCmd.AddCommand(registered_limit_setCmd)
}
