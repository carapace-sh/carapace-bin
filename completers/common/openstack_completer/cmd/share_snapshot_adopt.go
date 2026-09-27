package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_snapshot_adoptCmd = &cobra.Command{
	Use:   "adopt",
	Short: "Adopt a share snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_snapshot_adoptCmd).Standalone()

	share_snapshot_adoptCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_snapshot_adoptCmd.Flags().String("description", "", "Optional snapshot description (Default=None).")
	share_snapshot_adoptCmd.Flags().String("driver-option", "", "Set driver options as key=value pairs.(repeat")
	share_snapshot_adoptCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_snapshot_adoptCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_snapshot_adoptCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_snapshot_adoptCmd.Flags().String("name", "", "Optional snapshot name (Default=None).")
	share_snapshot_adoptCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_snapshot_adoptCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_snapshot_adoptCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_snapshot_adoptCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_snapshot_adoptCmd.Flags().Bool("wait", false, "Wait until share snapshot is adopted")
	share_snapshotCmd.AddCommand(share_snapshot_adoptCmd)
}
