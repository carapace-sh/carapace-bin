package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var role_assignment_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List role assignments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(role_assignment_listCmd).Standalone()

	role_assignment_listCmd.Flags().Bool("auth-project", false, "Only list assignments for the project to which the authenticated user's token is scoped")
	role_assignment_listCmd.Flags().Bool("auth-user", false, "Only list assignments for the authenticated user")
	role_assignment_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	role_assignment_listCmd.Flags().String("domain", "", "Domain to filter (name or ID)")
	role_assignment_listCmd.Flags().Bool("effective", false, "Returns only effective role assignments")
	role_assignment_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	role_assignment_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	role_assignment_listCmd.Flags().String("group", "", "Group to filter (name or ID)")
	role_assignment_listCmd.Flags().String("group-domain", "", "Domain the group belongs to (name or ID).")
	role_assignment_listCmd.Flags().Bool("inherited", false, "Specifies if the role grant is inheritable to the sub projects")
	role_assignment_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	role_assignment_listCmd.Flags().Bool("names", false, "Display names instead of IDs")
	role_assignment_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	role_assignment_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	role_assignment_listCmd.Flags().String("project", "", "Project to filter (name or ID)")
	role_assignment_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	role_assignment_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	role_assignment_listCmd.Flags().String("role", "", "Role to filter (name or ID)")
	role_assignment_listCmd.Flags().String("role-domain", "", "Domain the role belongs to (name or ID).")
	role_assignment_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	role_assignment_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	role_assignment_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	role_assignment_listCmd.Flags().String("system", "", "Filter based on system role assignments")
	role_assignment_listCmd.Flags().String("user", "", "User to filter (name or ID)")
	role_assignment_listCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	role_assignmentCmd.AddCommand(role_assignment_listCmd)
}
