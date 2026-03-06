package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:   "history <subcommand> [options] [<transaction-spec>]",
	Short: "manage transaction history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var historyListCmd = &cobra.Command{
	Use:   "list [transaction-spec]...",
	Short: "list transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var historyInfoCmd = &cobra.Command{
	Use:   "info [transaction-spec]...",
	Short: "print details about transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var historyUndoCmd = &cobra.Command{
	Use:   "undo <transaction-spec>",
	Short: "revert all actions from the specified transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var historyRedoCmd = &cobra.Command{
	Use:   "redo <transaction-spec>",
	Short: "repeat the specified transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var historyRollbackCmd = &cobra.Command{
	Use:   "rollback <transaction-spec>",
	Short: "undo all transactions after the specified transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var historyStoreCmd = &cobra.Command{
	Use:   "store [transaction-spec]",
	Short: "store the transaction into a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyCmd).Standalone()

	historyCmd.AddCommand(historyListCmd)
	historyCmd.AddCommand(historyInfoCmd)
	historyCmd.AddCommand(historyUndoCmd)
	historyCmd.AddCommand(historyRedoCmd)
	historyCmd.AddCommand(historyRollbackCmd)
	historyCmd.AddCommand(historyStoreCmd)

	for _, subcmd := range []*cobra.Command{historyListCmd, historyInfoCmd} {
		subcmd.Flags().Bool("reverse", false, "reverse the order of transactions")
		subcmd.Flags().String("contains-pkgs", "", "show only transactions containing packages")
		subcmd.Flags().Bool("json", false, "request JSON output format")
	}

	for _, subcmd := range []*cobra.Command{historyUndoCmd, historyRollbackCmd, historyRedoCmd} {
		subcmd.Flags().Bool("skip-unavailable", false, "allow skipping packages")
	}

	for _, subcmd := range []*cobra.Command{historyUndoCmd, historyRollbackCmd} {
		subcmd.Flags().Bool("ignore-extras", false, "don't consider extra packages as errors")
		subcmd.Flags().Bool("ignore-installed", false, "don't consider mismatches as errors")
	}

	rootCmd.AddCommand(historyCmd)
}
