package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var consistency_group_snapshot_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new consistency group snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consistency_group_snapshot_createCmd).Standalone()

	consistency_group_snapshot_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	consistency_group_snapshot_createCmd.Flags().String("consistency-group", "", "Consistency group to snapshot (name or ID) (default to be the same as <snapshot-name>)")
	consistency_group_snapshot_createCmd.Flags().String("description", "", "Description of this consistency group snapshot")
	consistency_group_snapshot_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	consistency_group_snapshot_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	consistency_group_snapshot_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	consistency_group_snapshot_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	consistency_group_snapshot_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	consistency_group_snapshot_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	consistency_group_snapshot_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	consistency_group_snapshotCmd.AddCommand(consistency_group_snapshot_createCmd)
}
