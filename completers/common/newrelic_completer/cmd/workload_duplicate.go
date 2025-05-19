package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workload_duplicateCmd = &cobra.Command{
	Use:   "duplicate",
	Short: "Duplicate a New Relic One workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workload_duplicateCmd).Standalone()
	workload_duplicateCmd.Flags().StringP("guid", "g", "", "the GUID of the workload you want to duplicate")
	workload_duplicateCmd.Flags().StringP("name", "n", "", "the name of the workload to duplicate")
	workloadCmd.AddCommand(workload_duplicateCmd)
}
