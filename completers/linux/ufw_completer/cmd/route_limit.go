package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route_limitCmd = &cobra.Command{
	Use:   "limit",
	Short: "add route limit rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_limitCmd).Standalone()

	routeCmd.AddCommand(route_limitCmd)
}
