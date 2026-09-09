package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tap_service_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a tap service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tap_service_deleteCmd).Standalone()

	tap_serviceCmd.AddCommand(tap_service_deleteCmd)
}
