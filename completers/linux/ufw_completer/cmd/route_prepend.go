package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route_prependCmd = &cobra.Command{
	Use:   "prepend",
	Short: "prepend route RULE",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_prependCmd).Standalone()

	routeCmd.AddCommand(route_prependCmd)
}
