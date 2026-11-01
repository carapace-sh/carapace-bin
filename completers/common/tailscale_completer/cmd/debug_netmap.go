package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_netmapCmd = &cobra.Command{
	Use:   "netmap",
	Short: "Print the current network map",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_netmapCmd).Standalone()

	debugCmd.AddCommand(debug_netmapCmd)
}
