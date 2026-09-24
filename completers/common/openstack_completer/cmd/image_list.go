package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_listCmd).Standalone()

	image_listCmd.Flags().Bool("all", false, "List all images")
	image_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	image_listCmd.Flags().Bool("community", false, "List only community images (requires --os-image-api-version 2.5 or later)")
	image_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	image_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	image_listCmd.Flags().Bool("hidden", false, "List hidden images")
	image_listCmd.Flags().String("limit", "", "The maximum number of entries to return per page.")
	image_listCmd.Flags().Bool("long", false, "List additional fields in output")
	image_listCmd.Flags().String("marker", "", "The first position in the collection to return results from.")
	image_listCmd.Flags().String("max-items", "", "The maximum number of entries to return in total, paging through multiple requests if needed.")
	image_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	image_listCmd.Flags().String("member-status", "", "Filter images based on member status.")
	image_listCmd.Flags().String("name", "", "Filter images based on name.")
	image_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	image_listCmd.Flags().String("page-size", "", "==SUPPRESS==")
	image_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	image_listCmd.Flags().Bool("private", false, "List only private images")
	image_listCmd.Flags().String("project", "", "Search by project (admin only) (name or ID)")
	image_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	image_listCmd.Flags().String("property", "", "Filter output based on property (repeat option to filter on multiple properties)")
	image_listCmd.Flags().Bool("public", false, "List only public images")
	image_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	image_listCmd.Flags().Bool("shared", false, "List only shared images (requires --os-image-api-version 2.5 or later)")
	image_listCmd.Flags().String("sort", "", "Sort output by selected keys and directions (asc or desc) (default: name:asc), multiple keys and directions can be specified separated by comma")
	image_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	image_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	image_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	image_listCmd.Flags().String("status", "", "Filter images based on status.")
	image_listCmd.Flags().String("tag", "", "Filter images based on tag.")
	imageCmd.AddCommand(image_listCmd)
}
