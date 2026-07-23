package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route_replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "change or add a new route",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_replaceCmd).Standalone()

	routeCmd.AddCommand(route_replaceCmd)
}
