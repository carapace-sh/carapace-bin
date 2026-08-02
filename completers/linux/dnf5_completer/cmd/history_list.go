package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var historyListCmd = &cobra.Command{
	Use:   "list [options] [<transaction-spec>...]",
	Short: "list transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyListCmd).Standalone()

	historyListCmd.Flags().String("contains-pkgs", "", "Filter by package names")
	historyListCmd.Flags().Bool("json", false, "Request json output format")
	historyListCmd.Flags().Bool("reverse", false, "Reverse order")

	historyCmd.AddCommand(historyListCmd)
}
