package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hypervisor_statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hypervisor_statsCmd).Standalone()

	hypervisorCmd.AddCommand(hypervisor_statsCmd)
}
