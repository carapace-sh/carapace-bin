package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_auto_allocatedCmd = &cobra.Command{
	Use:   "allocated",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_auto_allocatedCmd).Standalone()

	network_autoCmd.AddCommand(network_auto_allocatedCmd)
}
