package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_listClientsCmd = &cobra.Command{
	Use:   "list-clients",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_listClientsCmd).Standalone()

	action_helpCmd.AddCommand(action_help_listClientsCmd)
}
