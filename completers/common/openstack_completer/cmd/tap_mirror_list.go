package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tap_mirror_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tap mirrors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tap_mirror_listCmd).Standalone()

	tap_mirror_listCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	tap_mirror_listCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	tap_mirror_listCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	tap_mirror_listCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	tap_mirror_listCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	tap_mirror_listCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	tap_mirror_listCmd.Flags().String("project", "", "Owner's project (name or ID)")
	tap_mirror_listCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	tap_mirror_listCmd.Flags().String("quote", "", "when to include quotes, defaults to nonnumeric")
	tap_mirror_listCmd.Flags().Bool("sort-ascending", false, "sort the column(s) in ascending order")
	tap_mirror_listCmd.Flags().String("sort-column", "", "specify the column(s) to sort the data (columns specified first have a priority, non-existing columns are ignored), can be repeated")
	tap_mirror_listCmd.Flags().Bool("sort-descending", false, "sort the column(s) in descending order")
	tap_mirrorCmd.AddCommand(tap_mirror_listCmd)
}
