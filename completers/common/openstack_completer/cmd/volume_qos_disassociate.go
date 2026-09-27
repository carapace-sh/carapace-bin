package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_qos_disassociateCmd = &cobra.Command{
	Use:   "disassociate",
	Short: "Disassociate a QoS specification from a volume type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_qos_disassociateCmd).Standalone()

	volume_qos_disassociateCmd.Flags().Bool("all", false, "Disassociate the QoS from every volume type")
	volume_qos_disassociateCmd.Flags().String("volume-type", "", "Volume type to disassociate the QoS from (name or ID)")
	volume_qosCmd.AddCommand(volume_qos_disassociateCmd)
}
