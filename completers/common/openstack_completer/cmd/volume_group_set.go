package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_group_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update a volume group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_group_setCmd).Standalone()

	volume_group_setCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_group_setCmd.Flags().String("description", "", "New description for group.")
	volume_group_setCmd.Flags().Bool("disable-replication", false, "Disable replication for group.")
	volume_group_setCmd.Flags().Bool("enable-replication", false, "Enable replication for group.")
	volume_group_setCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_group_setCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_group_setCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_group_setCmd.Flags().String("name", "", "New name for group.")
	volume_group_setCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_group_setCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_group_setCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_group_setCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_groupCmd.AddCommand(volume_group_setCmd)
}
