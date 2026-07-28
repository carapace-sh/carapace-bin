package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var routeCmd = &cobra.Command{
	Use:     "route",
	Aliases: []string{"r"},
	Short:   "routing table entry",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(routeCmd).Standalone()

	rootCmd.AddCommand(routeCmd)
}
