package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_qos_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set QoS specification properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_qos_setCmd).Standalone()

	volume_qos_setCmd.Flags().Bool("no-property", false, "Remove all properties from <qos-spec> (specify both --no-property and --property to remove the current properties before setting new properties)")
	volume_qos_setCmd.Flags().String("property", "", "Property to add or modify for this QoS specification (repeat option to set multiple properties)")
	volume_qosCmd.AddCommand(volume_qos_setCmd)
}
