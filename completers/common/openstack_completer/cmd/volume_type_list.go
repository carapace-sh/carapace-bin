package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_type_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List volume types",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_type_listCmd).Standalone()

	volume_type_listCmd.Flags().String("availability-zone", "", "List only volume types with this availability configured (this is an alias for '--property RESKEY:availability_zones:<az>') (repeat option to filter on multiple availability zones)")
	volume_type_listCmd.Flags().Bool("cacheable", false, "List only volume types with caching enabled (this is an alias for '--property cacheable=<is> True') (admin only) (supported by --os-volume-api-version 3.52 or above)")
	volume_type_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_type_listCmd.Flags().Bool("default", false, "List the default volume type")
	volume_type_listCmd.Flags().Bool("encryption-type", false, "Display encryption information for each volume type (admin only)")
	volume_type_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_type_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_type_listCmd.Flags().Bool("long", false, "List additional fields in output")
	volume_type_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_type_listCmd.Flags().Bool("multiattach", false, "List only volume types with multi-attach enabled (this is an alias for '--property multiattach=<is> True') (supported by --os-volume-api-version 3.52 or above)")
	volume_type_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_type_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_type_listCmd.Flags().Bool("private", false, "List only private types (admin only)")
	volume_type_listCmd.Flags().String("property", "", "Filter by a property on the volume types (repeat option to filter by multiple properties) (admin only except for user-visible extra specs) (supported by --os-volume-api-version 3.52 or above)")
	volume_type_listCmd.Flags().Bool("public", false, "List only public types")
	volume_type_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	volume_type_listCmd.Flags().Bool("replicated", false, "List only volume types with replication enabled (this is an alias for '--property replication_enabled=<is> True') (supported by --os-volume-api-version 3.52 or above)")
	volume_type_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	volume_type_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	volume_type_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	volume_typeCmd.AddCommand(volume_type_listCmd)
}
