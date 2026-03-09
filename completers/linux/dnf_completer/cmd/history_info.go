package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_infoCmd = &cobra.Command{
	Use:   "info [transaction-spec]...",
	Short: "print details about transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_infoCmd).Standalone()

	history_infoCmd.Flags().String("contains-pkgs", "", "show only transactions containing packages")
	history_infoCmd.Flags().Bool("json", false, "request JSON output format")
	history_infoCmd.Flags().Bool("reverse", false, "reverse the order of transactions")
	historyCmd.AddCommand(history_infoCmd)
}
