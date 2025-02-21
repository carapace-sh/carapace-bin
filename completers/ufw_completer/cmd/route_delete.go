package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete route RULE",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_deleteCmd).Standalone()

	routeCmd.AddCommand(route_deleteCmd)
}
