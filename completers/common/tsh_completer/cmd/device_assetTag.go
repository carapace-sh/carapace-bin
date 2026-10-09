package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var device_assetTagCmd = &cobra.Command{
	Use:    "asset-tag",
	Short:  "Print the detected device asset tag.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(device_assetTagCmd).Standalone()

	deviceCmd.AddCommand(device_assetTagCmd)
}
