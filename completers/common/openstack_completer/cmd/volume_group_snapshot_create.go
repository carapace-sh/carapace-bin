package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_group_snapshot_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a volume group snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_group_snapshot_createCmd).Standalone()

	volume_group_snapshot_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_group_snapshot_createCmd.Flags().String("description", "", "Description of a volume group snapshot.")
	volume_group_snapshot_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_group_snapshot_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_group_snapshot_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_group_snapshot_createCmd.Flags().String("name", "", "Name of the volume group snapshot.")
	volume_group_snapshot_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_group_snapshot_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_group_snapshot_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_group_snapshot_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_group_snapshotCmd.AddCommand(volume_group_snapshot_createCmd)
}
