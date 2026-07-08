package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route_changeCmd = &cobra.Command{
	Use:   "change",
	Short: "change a route",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_changeCmd).Standalone()

	routeCmd.AddCommand(route_changeCmd)
}
