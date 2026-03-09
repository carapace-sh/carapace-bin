package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_listCmd = &cobra.Command{
	Use:   "list [transaction-spec]...",
	Short: "list transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_listCmd).Standalone()

	history_listCmd.Flags().String("contains-pkgs", "", "show only transactions containing packages")
	history_listCmd.Flags().Bool("json", false, "request JSON output format")
	history_listCmd.Flags().Bool("reverse", false, "reverse the order of transactions")
	historyCmd.AddCommand(history_listCmd)
}
