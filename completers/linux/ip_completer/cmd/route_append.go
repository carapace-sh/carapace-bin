package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route_appendCmd = &cobra.Command{
	Use:   "append",
	Short: "append a route",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_appendCmd).Standalone()

	routeCmd.AddCommand(route_appendCmd)
}
