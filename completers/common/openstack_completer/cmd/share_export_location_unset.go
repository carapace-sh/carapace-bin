package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_export_location_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset a share export location property",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_export_location_unsetCmd).Standalone()

	share_export_location_unsetCmd.Flags().String("property", "", "Remove a property from export location (repeat option to remove multiple properties).")
	share_export_locationCmd.AddCommand(share_export_location_unsetCmd)
}
