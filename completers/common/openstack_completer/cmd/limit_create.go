package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var limit_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a limit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(limit_createCmd).Standalone()

	limit_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	limit_createCmd.Flags().String("description", "", "Description of the limit")
	limit_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	limit_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	limit_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	limit_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	limit_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	limit_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	limit_createCmd.Flags().String("project", "", "Project to associate the resource limit to")
	limit_createCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	limit_createCmd.Flags().String("region", "", "Region for the limit to affect.")
	limit_createCmd.Flags().String("resource-limit", "", "The resource limit for the project to assume")
	limit_createCmd.Flags().String("service", "", "Service responsible for the resource to limit")
	limit_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	limit_createCmd.MarkFlagRequired("project")
	limit_createCmd.MarkFlagRequired("resource-limit")
	limit_createCmd.MarkFlagRequired("service")
	limitCmd.AddCommand(limit_createCmd)
}
