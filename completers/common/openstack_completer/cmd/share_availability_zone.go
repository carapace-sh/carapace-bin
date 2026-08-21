package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var share_availability_zoneCmd = &cobra.Command{
	Use:   "zone",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(share_availability_zoneCmd).Standalone()

	share_availabilityCmd.AddCommand(share_availability_zoneCmd)
}
