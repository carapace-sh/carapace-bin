package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var clientInspectCmd = &cobra.Command{
	Use:   "inspect [CLIENT]",
	Short: "Inspect a client",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clientInspectCmd).Standalone()

	clientInspectCmd.Flags().Bool("json", false, "output the client and its actions as JSON")

	clientCmd.AddCommand(clientInspectCmd)

	carapace.Gen(clientInspectCmd).PositionalCompletion(
		action.ActionClients(clientInspectCmd),
	)
}
