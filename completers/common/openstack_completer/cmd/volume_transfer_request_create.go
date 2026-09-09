package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_transfer_request_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create volume transfer request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_transfer_request_createCmd).Standalone()

	volume_transfer_request_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_transfer_request_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_transfer_request_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_transfer_request_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_transfer_request_createCmd.Flags().String("name", "", "New transfer request name (default to None)")
	volume_transfer_request_createCmd.Flags().Bool("no-snapshots", false, "Disallow transfer volumes without snapshots (supported by --os-volume-api-version 3.55 or later)")
	volume_transfer_request_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_transfer_request_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_transfer_request_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_transfer_request_createCmd.Flags().Bool("snapshots", false, "Allow transfer volumes without snapshots (default) (supported by --os-volume-api-version 3.55 or later)")
	volume_transfer_request_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_transfer_requestCmd.AddCommand(volume_transfer_request_createCmd)
}
