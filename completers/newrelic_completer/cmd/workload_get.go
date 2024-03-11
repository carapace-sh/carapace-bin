package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workload_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a New Relic One workload.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workload_getCmd).Standalone()
	workload_getCmd.Flags().StringP("guid", "g", "", "the GUID of the workload")
	workloadCmd.AddCommand(workload_getCmd)
}
