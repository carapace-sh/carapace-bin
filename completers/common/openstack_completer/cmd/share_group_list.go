package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List share groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_listCmd).Standalone()

	share_group_listCmd.Flags().Bool("all-projects", false, "Display share groups from all projects (Admin only).")
	share_group_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_group_listCmd.Flags().String("description", "", "Filter results by description.")
	share_group_listCmd.Flags().String("description~", "", "Filter results matching a share group description pattern.")
	share_group_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_group_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_group_listCmd.Flags().String("host", "", "Filter results by host.")
	share_group_listCmd.Flags().String("limit", "", "Limit the number of share groups returned")
	share_group_listCmd.Flags().String("marker", "", "The last share group ID of the previous page")
	share_group_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_group_listCmd.Flags().String("name", "", "Filter results by name.")
	share_group_listCmd.Flags().String("name~", "", "Filter results matching a share group name pattern.")
	share_group_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_group_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_group_listCmd.Flags().String("project", "", "Filter results by project name or ID.")
	share_group_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	share_group_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_group_listCmd.Flags().String("share-group-type", "", "Filter results by a share group type ID or name that was used for share group creation. ")
	share_group_listCmd.Flags().String("share-network", "", "Filter results by share-network name or ID. ")
	share_group_listCmd.Flags().String("share-server", "", "Filter results by share server ID.")
	share_group_listCmd.Flags().String("snapshot", "", "Filter results by share group snapshot name or ID that was used to create the share group. ")
	share_group_listCmd.Flags().String("sort", "", "Sort output by selected keys and directions(asc or desc) (default: name:asc), multiple keys and directions can be specified separated by comma")
	share_group_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_group_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_group_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_group_listCmd.Flags().String("status", "", "Filter results by status.")
	share_groupCmd.AddCommand(share_group_listCmd)
}
