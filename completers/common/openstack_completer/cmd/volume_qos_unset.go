package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_qos_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset QoS specification properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_qos_unsetCmd).Standalone()

	volume_qos_unsetCmd.Flags().String("property", "", "Property to remove from the QoS specification.")
	volume_qosCmd.AddCommand(volume_qos_unsetCmd)
}
