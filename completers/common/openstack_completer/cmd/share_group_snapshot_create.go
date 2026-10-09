package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_snapshot_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a share group snapshot of the given share group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_snapshot_createCmd).Standalone()

	share_group_snapshot_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_group_snapshot_createCmd.Flags().String("description", "", "Optional share group snapshot description.")
	share_group_snapshot_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_group_snapshot_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_group_snapshot_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_group_snapshot_createCmd.Flags().String("name", "", "Optional share group snapshot name.")
	share_group_snapshot_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_group_snapshot_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_group_snapshot_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_group_snapshot_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_group_snapshot_createCmd.Flags().Bool("wait", false, "Wait for share group snapshot creation")
	share_group_snapshotCmd.AddCommand(share_group_snapshot_createCmd)
}
