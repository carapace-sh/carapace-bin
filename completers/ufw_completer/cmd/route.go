package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var routeCmd = &cobra.Command{
	Use:   "route",
	Short: "add route RULE",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(routeCmd).Standalone()

	rootCmd.AddCommand(routeCmd)
}
