package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/nmcli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var connection_downCmd = &cobra.Command{
	Use:   "down",
	Short: "Deactivate a connection",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_downCmd).Standalone()

	connectionCmd.AddCommand(connection_downCmd)

	carapace.Gen(connection_downCmd).PositionalCompletion(
		action.ActionConnections(),
	)
}
