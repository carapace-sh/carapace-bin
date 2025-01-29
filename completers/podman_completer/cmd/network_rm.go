package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_rmCmd = &cobra.Command{
	Use:     "rm [options] NETWORK [NETWORK...]",
	Short:   "Remove networks",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_rmCmd).Standalone()

	network_rmCmd.Flags().BoolP("force", "f", false, "remove any containers using network")
	network_rmCmd.Flags().StringP("time", "t", "", "Seconds to wait for running containers to stop before killing the container")
	networkCmd.AddCommand(network_rmCmd)
}
