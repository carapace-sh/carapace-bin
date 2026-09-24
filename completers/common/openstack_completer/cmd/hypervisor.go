package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hypervisorCmd = &cobra.Command{
	Use:   "hypervisor",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hypervisorCmd).Standalone()

	rootCmd.AddCommand(hypervisorCmd)
}
