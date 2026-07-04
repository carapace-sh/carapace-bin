package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route_denyCmd = &cobra.Command{
	Use:   "deny",
	Short: "add route deny rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_denyCmd).Standalone()

	routeCmd.AddCommand(route_denyCmd)
}
