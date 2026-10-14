package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var availability_zoneCmd = &cobra.Command{
	Use:   "zone",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(availability_zoneCmd).Standalone()

	availabilityCmd.AddCommand(availability_zoneCmd)
}
