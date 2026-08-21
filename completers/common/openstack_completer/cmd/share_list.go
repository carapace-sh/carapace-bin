package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List shares",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_listCmd).Standalone()

	share_listCmd.Flags().Bool("all-projects", false, "Include all projects (admin only)")
	share_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_listCmd.Flags().String("description~", "", "Filter results matching a share description pattern.Available only for microversion >= 2.36.")
	share_listCmd.Flags().String("encryption-key-ref", "", "Filter shares by their encryption key ref.")
	share_listCmd.Flags().String("export-location", "", "Filter shares by export location id or path.")
	share_listCmd.Flags().String("extra-spec", "", "Filter shares with extra specs (key=value) of the share type that they belong to.")
	share_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_listCmd.Flags().String("host", "", "Filter shares belonging to a given host (admin only)")
	share_listCmd.Flags().String("limit", "", "Maximum number of shares to display")
	share_listCmd.Flags().Bool("long", false, "List additional fields in output")
	share_listCmd.Flags().String("marker", "", "The last share ID of the previous page")
	share_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_listCmd.Flags().String("name", "", "Filter shares by share name")
	share_listCmd.Flags().String("name~", "", "Filter results matching a share name pattern.")
	share_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_listCmd.Flags().String("project", "", "Filter shares by project (name or ID) (admin only)")
	share_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	share_listCmd.Flags().String("property", "", "Filter shares having a given metadata key=value property (repeat option to filter by multiple properties)")
	share_listCmd.Flags().Bool("public", false, "Include public shares")
	share_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	share_listCmd.Flags().String("share-group", "", "Filter shares belonging to a given share group")
	share_listCmd.Flags().String("share-network", "", "Filter shares exported on a given share network")
	share_listCmd.Flags().String("share-server", "", "Filter shares exported via a given share server (admin only)")
	share_listCmd.Flags().String("share-type", "", "Filter shares of a given share type")
	share_listCmd.Flags().String("snapshot", "", "Filter shares by snapshot name or id.")
	share_listCmd.Flags().Bool("soft-deleted", false, "Get shares in recycle bin.")
	share_listCmd.Flags().String("sort", "", "Sort output by selected keys and directions(asc or desc) (default: name:asc), multiple keys and directions can be specified separated by comma")
	share_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	share_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	share_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	share_listCmd.Flags().String("status", "", "Filter shares by status")
	share_listCmd.Flags().String("user", "", "Filter results by user (name or ID) (admin only)")
	share_listCmd.Flags().String("user-domain", "", "Domain the user belongs to (name or ID).")
	shareCmd.AddCommand(share_listCmd)
}
