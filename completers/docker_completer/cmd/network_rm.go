package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var network_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more networks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_rmCmd).Standalone()

	networkCmd.AddCommand(network_rmCmd)

	carapace.Gen(network_rmCmd).PositionalAnyCompletion(action.ActionNetworks())
}
