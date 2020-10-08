package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/nmcli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var connection_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing connection profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_deleteCmd).Standalone()

	connectionCmd.AddCommand(connection_deleteCmd)

	carapace.Gen(connection_deleteCmd).PositionalCompletion(
		action.ActionConnections(),
	)
}
