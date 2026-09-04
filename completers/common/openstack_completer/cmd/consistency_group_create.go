package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var consistency_group_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new consistency group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consistency_group_createCmd).Standalone()

	consistency_group_createCmd.Flags().String("availability-zone", "", "Availability zone for this consistency group (not available if creating consistency group from source)")
	consistency_group_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	consistency_group_createCmd.Flags().String("consistency-group-snapshot", "", "==SUPPRESS==")
	consistency_group_createCmd.Flags().String("consistency-group-source", "", "==SUPPRESS==")
	consistency_group_createCmd.Flags().String("description", "", "Description of this consistency group")
	consistency_group_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	consistency_group_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	consistency_group_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	consistency_group_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	consistency_group_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	consistency_group_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	consistency_group_createCmd.Flags().String("snapshot", "", "Existing consistency group snapshot (name or ID)")
	consistency_group_createCmd.Flags().String("source", "", "Existing consistency group (name or ID)")
	consistency_group_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	consistency_group_createCmd.Flags().String("volume-type", "", "Volume type of this consistency group (name or ID)")
	consistency_groupCmd.AddCommand(consistency_group_createCmd)
}
