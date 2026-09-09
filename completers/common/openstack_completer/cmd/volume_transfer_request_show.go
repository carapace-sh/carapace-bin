package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_transfer_request_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show volume transfer request details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_transfer_request_showCmd).Standalone()

	volume_transfer_request_showCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_transfer_request_showCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_transfer_request_showCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_transfer_request_showCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_transfer_request_showCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_transfer_request_showCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_transfer_request_showCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_transfer_request_showCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_transfer_requestCmd.AddCommand(volume_transfer_request_showCmd)
}
