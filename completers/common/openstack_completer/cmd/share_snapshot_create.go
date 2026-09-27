package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_snapshot_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a snapshot of the given share",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_snapshot_createCmd).Standalone()

	share_snapshot_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_snapshot_createCmd.Flags().String("description", "", "Add a description to the snapshot (Optional).")
	share_snapshot_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_snapshot_createCmd.Flags().Bool("force", false, "Optional flag to indicate whether to snapshot a share even if it's busy.")
	share_snapshot_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_snapshot_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_snapshot_createCmd.Flags().String("name", "", "Add a name to the snapshot (Optional).")
	share_snapshot_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_snapshot_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_snapshot_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_snapshot_createCmd.Flags().String("property", "", "Set a property to this snapshot (repeat option to set multiple properties).Available only for microversion >= 2.73")
	share_snapshot_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_snapshot_createCmd.Flags().Bool("wait", false, "Wait for share snapshot creation")
	share_snapshotCmd.AddCommand(share_snapshot_createCmd)
}
