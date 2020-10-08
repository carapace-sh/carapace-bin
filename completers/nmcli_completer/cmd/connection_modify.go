package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/nmcli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var connection_modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify one or more properties of the connection profile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connection_modifyCmd).Standalone()

	connection_modifyCmd.Flags().Bool("temporary", false, "only temporary")
	connectionCmd.AddCommand(connection_modifyCmd)

	carapace.Gen(connection_modifyCmd).PositionalCompletion(
		action.ActionConnections(),
	)
}
