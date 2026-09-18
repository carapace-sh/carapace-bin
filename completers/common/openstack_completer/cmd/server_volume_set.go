package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_volume_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update a volume attachment on the server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_volume_setCmd).Standalone()

	server_volume_setCmd.Flags().Bool("delete-on-termination", false, "Delete the volume when the server is destroyed (supported by --os-compute-api-version 2.85 or above)")
	server_volume_setCmd.Flags().Bool("preserve-on-termination", false, "Preserve the volume when the server is destroyed (supported by --os-compute-api-version 2.85 or above)")
	server_volumeCmd.AddCommand(server_volume_setCmd)
}
