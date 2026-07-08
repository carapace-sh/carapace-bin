package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route_saveCmd = &cobra.Command{
	Use:   "save",
	Short: "save routing table to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_saveCmd).Standalone()

	routeCmd.AddCommand(route_saveCmd)
}
