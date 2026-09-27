package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_transfer_request_acceptCmd = &cobra.Command{
	Use:   "accept",
	Short: "Accept volume transfer request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_transfer_request_acceptCmd).Standalone()

	volume_transfer_request_acceptCmd.Flags().String("auth-key", "", "Volume transfer request authentication key")
	volume_transfer_request_acceptCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	volume_transfer_request_acceptCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	volume_transfer_request_acceptCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	volume_transfer_request_acceptCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	volume_transfer_request_acceptCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	volume_transfer_request_acceptCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	volume_transfer_request_acceptCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	volume_transfer_request_acceptCmd.Flags().String("variable", "", "==SUPPRESS==")
	volume_transfer_request_acceptCmd.MarkFlagRequired("auth-key")
	volume_transfer_requestCmd.AddCommand(volume_transfer_request_acceptCmd)
}
