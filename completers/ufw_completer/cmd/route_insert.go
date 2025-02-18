package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route_insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "insert route RULE at NUM",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route_insertCmd)

	routeCmd.AddCommand(route_insertCmd)
}
