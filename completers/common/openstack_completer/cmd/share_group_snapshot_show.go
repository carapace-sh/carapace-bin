package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_group_snapshot_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show details about a share group snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_group_snapshot_showCmd).Standalone()

	share_group_snapshot_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_group_snapshot_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_group_snapshot_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_group_snapshot_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_group_snapshot_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_group_snapshot_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_group_snapshot_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_group_snapshot_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_group_snapshotCmd.AddCommand(share_group_snapshot_showCmd)
}
