package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workloadCmd = &cobra.Command{
	Use:   "workload",
	Short: "Interact with New Relic One workloads",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workloadCmd).Standalone()
	rootCmd.AddCommand(workloadCmd)
}
