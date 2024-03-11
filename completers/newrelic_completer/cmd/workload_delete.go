package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workload_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a New Relic One workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workload_deleteCmd).Standalone()
	workload_deleteCmd.Flags().StringP("guid", "g", "", "the GUID of the workload to delete")
	workloadCmd.AddCommand(workload_deleteCmd)
}
