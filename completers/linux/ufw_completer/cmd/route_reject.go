package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route_rejectCmd = &cobra.Command{
	Use:   "reject",
	Short: "add route reject rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_rejectCmd).Standalone()

	routeCmd.AddCommand(route_rejectCmd)
}
