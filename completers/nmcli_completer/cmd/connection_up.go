package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/nmcli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var connection_upCmd = &cobra.Command{
	Use:   "up",
	Short: "Activate a connection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_upCmd).Standalone()

	connectionCmd.AddCommand(connection_upCmd)

	carapace.Gen(connection_upCmd).PositionalCompletion(
		action.ActionConnections(),
	)
}
