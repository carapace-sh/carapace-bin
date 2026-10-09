package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_qos_type_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set qos type description or specs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_qos_type_setCmd).Standalone()

	share_qos_type_setCmd.Flags().String("description", "", "New description of qos type.")
	share_qos_type_setCmd.Flags().String("spec", "", "Spec key and value of qos type that will be used for QoS type.")
	share_qos_typeCmd.AddCommand(share_qos_type_setCmd)
}
