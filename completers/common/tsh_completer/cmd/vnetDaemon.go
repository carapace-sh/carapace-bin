package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vnetDaemonCmd = &cobra.Command{
	Use:    "vnet-daemon",
	Short:  "Start the VNet D-Bus daemon.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vnetDaemonCmd).Standalone()

	rootCmd.AddCommand(vnetDaemonCmd)
}
