package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nrql_historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Retrieve NRQL query history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nrql_historyCmd).Standalone()
	nrql_historyCmd.Flags().IntP("limit", "l", 10, "history items to return (default: 10, max: 100)")
	nrqlCmd.AddCommand(nrql_historyCmd)
}
