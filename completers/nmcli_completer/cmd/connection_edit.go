package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/nmcli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var connection_editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit an existing connection profile in an interactive editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_editCmd).Standalone()

	connectionCmd.AddCommand(connection_editCmd)

	carapace.Gen(connection_editCmd).PositionalCompletion(
		action.ActionConnections(),
	)
}
