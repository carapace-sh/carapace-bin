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

func init() {
	carapace.Gen(historyCmd).Standalone()

	rootCmd.AddCommand(historyCmd)
}
