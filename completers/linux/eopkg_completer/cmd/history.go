package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var historyCmd = &cobra.Command{
	Use:     "history",
	Aliases: []string{"hs"},
	Short:   "manage the eopkg transaction history",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyCmd).Standalone()

	historyCmd.Flags().IntP("last", "l", 0, "only output the last <n> operations")
	historyCmd.Flags().BoolP("snapshot", "s", false, "create a new snapshot transaction to record the current system state")
	historyCmd.Flags().StringP("takeback", "t", "", "roll the system state back to the state of the given transaction")

	rootCmd.AddCommand(historyCmd)
}
