package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_snapshot_access_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Allow access to a snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_snapshot_access_createCmd).Standalone()

	share_snapshot_access_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	share_snapshot_access_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	share_snapshot_access_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	share_snapshot_access_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	share_snapshot_access_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	share_snapshot_access_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	share_snapshot_access_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	share_snapshot_access_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	share_snapshot_accessCmd.AddCommand(share_snapshot_access_createCmd)
}
