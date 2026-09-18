package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tap_flow_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a tap flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tap_flow_deleteCmd).Standalone()

	tap_flowCmd.AddCommand(tap_flow_deleteCmd)
}
