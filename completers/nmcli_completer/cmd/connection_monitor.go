package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/nmcli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var connection_monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor an existing connection profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_monitorCmd).Standalone()

	connectionCmd.AddCommand(connection_monitorCmd)

	carapace.Gen(connection_monitorCmd).PositionalCompletion(
		action.ActionConnections(),
	)
}
