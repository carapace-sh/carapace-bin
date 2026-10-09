package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_snapshot_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new volume snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_snapshot_createCmd).Standalone()

	volume_snapshot_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_snapshot_createCmd.Flags().String("description", "", "Description of the snapshot")
	volume_snapshot_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_snapshot_createCmd.Flags().Bool("force", false, "Allow snapshot of in-use (attached) volume.")
	volume_snapshot_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_snapshot_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_snapshot_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_snapshot_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_snapshot_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_snapshot_createCmd.Flags().String("property", "", "Set a property to this snapshot (repeat option to set multiple properties)")
	volume_snapshot_createCmd.Flags().String("remote-source", "", "The attribute(s) of the existing remote volume snapshot (admin required) (repeat option to specify multiple attributes) e.g.: '--remote-source source-name=test_name --remote-source source-id=test_id'")
	volume_snapshot_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_snapshot_createCmd.Flags().String("volume", "", "Volume to snapshot (name or ID) (default is <snapshot-name>)")
	volume_snapshotCmd.AddCommand(volume_snapshot_createCmd)
}
