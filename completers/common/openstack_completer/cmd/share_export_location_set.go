package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_export_location_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an export location property.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_export_location_setCmd).Standalone()

	share_export_location_setCmd.Flags().String("property", "", "Set a property to this export location (repeat option to set multiple properties).")
	share_export_locationCmd.AddCommand(share_export_location_setCmd)
}
