package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var network_rmCmd = &cobra.Command{
	Use:     "rm NETWORK [NETWORK...]",
	Short:   "Remove one or more networks",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_rmCmd).Standalone()

	network_rmCmd.Flags().BoolP("force", "f", false, "Do not error if the network does not exist")
	networkCmd.AddCommand(network_rmCmd)

	carapace.Gen(network_rmCmd).PositionalAnyCompletion(docker.ActionNetworks())
}
