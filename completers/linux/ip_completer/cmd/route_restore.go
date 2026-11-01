package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route_restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "restore routing table from stdin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_restoreCmd).Standalone()

	routeCmd.AddCommand(route_restoreCmd)
}
