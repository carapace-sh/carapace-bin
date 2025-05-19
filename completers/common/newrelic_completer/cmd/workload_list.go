package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workload_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the New Relic One workloads for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workload_listCmd).Standalone()
	workloadCmd.AddCommand(workload_listCmd)
}
