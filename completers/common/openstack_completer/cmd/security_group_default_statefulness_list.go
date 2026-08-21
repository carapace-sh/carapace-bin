package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_group_default_statefulness_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List security group default statefulness settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_group_default_statefulness_listCmd).Standalone()

	security_group_default_statefulness_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	security_group_default_statefulness_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	security_group_default_statefulness_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	security_group_default_statefulness_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	security_group_default_statefulness_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	security_group_default_statefulness_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	security_group_default_statefulness_listCmd.Flags().String("project", "", "List only settings for this project (name or ID)")
	security_group_default_statefulness_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	security_group_default_statefulness_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	security_group_default_statefulness_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	security_group_default_statefulness_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	security_group_default_statefulness_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	security_group_default_statefulnessCmd.AddCommand(security_group_default_statefulness_listCmd)
}
