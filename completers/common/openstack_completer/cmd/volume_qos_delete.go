package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var volume_qos_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete QoS specification",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_qos_deleteCmd).Standalone()

	volume_qos_deleteCmd.Flags().Bool("force", false, "Allow to delete in-use QoS specification(s)")
	volume_qosCmd.AddCommand(volume_qos_deleteCmd)
}
