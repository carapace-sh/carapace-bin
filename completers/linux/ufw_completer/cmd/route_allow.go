package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route_allowCmd = &cobra.Command{
	Use:   "allow",
	Short: "add route allow rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_allowCmd).Standalone()

	routeCmd.AddCommand(route_allowCmd)
}
