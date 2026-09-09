package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var consistency_group_snapshot_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display consistency group snapshot details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(consistency_group_snapshot_showCmd).Standalone()

	consistency_group_snapshot_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	consistency_group_snapshot_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	consistency_group_snapshot_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	consistency_group_snapshot_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	consistency_group_snapshot_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	consistency_group_snapshot_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	consistency_group_snapshot_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	consistency_group_snapshot_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	consistency_group_snapshotCmd.AddCommand(consistency_group_snapshot_showCmd)
}
