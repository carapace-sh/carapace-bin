package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var idl_fetchHistoricalCmd = &cobra.Command{
	Use:   "fetch-historical",
	Short: "Fetches historical IDL versions for the given program from a cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(idl_fetchHistoricalCmd).Standalone()

	idl_fetchHistoricalCmd.Flags().String("after", "", "Fetch IDL after this date (YYYY-MM-DD)")
	idl_fetchHistoricalCmd.Flags().String("authority", "", "Fetch authority-scoped PMP metadata account history for this authority")
	idl_fetchHistoricalCmd.Flags().String("before", "", "Fetch IDL before this date (YYYY-MM-DD)")
	idl_fetchHistoricalCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	idl_fetchHistoricalCmd.Flags().String("max-signatures", "1000", "Hard cap on signatures fetched per history source")
	idl_fetchHistoricalCmd.Flags().Bool("no-parallel", false, "Force sequential transaction fetches (equivalent to --rpc-workers 1)")
	idl_fetchHistoricalCmd.Flags().String("out-dir", "", "Output directory for fetched versions (defaults to the current directory)")
	idl_fetchHistoricalCmd.Flags().String("rpc-max-retries", "5", "Max retry attempts per transaction on 429/timeout errors")
	idl_fetchHistoricalCmd.Flags().String("rpc-retry-backoff-ms", "500", "Base backoff in milliseconds between retries (doubled each attempt)")
	idl_fetchHistoricalCmd.Flags().String("rpc-workers", "", "Max parallel RPC workers for transaction fetches")
	idl_fetchHistoricalCmd.Flags().String("slot", "", "Fetch IDL at specific slot")
	idl_fetchHistoricalCmd.Flags().Bool("verbose", false, "Print diagnostic progress messages")
	idlCmd.AddCommand(idl_fetchHistoricalCmd)

	carapace.Gen(idl_fetchHistoricalCmd).FlagCompletion(carapace.ActionMap{
		"out-dir": carapace.ActionFiles(),
	})
}
