package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var historyInfoCmd = &cobra.Command{
	Use:   "info [options] [<transaction-spec>...]",
	Short: "print details about transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(historyInfoCmd).Standalone()

	historyInfoCmd.Flags().String("contains-pkgs", "", "Filter by package names")
	historyInfoCmd.Flags().Bool("json", false, "Request json output format")
	historyInfoCmd.Flags().Bool("reverse", false, "Reverse order")

	historyCmd.AddCommand(historyInfoCmd)
}
