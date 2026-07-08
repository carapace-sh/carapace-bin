package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route_getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a single route",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_getCmd).Standalone()

	routeCmd.AddCommand(route_getCmd)
}
