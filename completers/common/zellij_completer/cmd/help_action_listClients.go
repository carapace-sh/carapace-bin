package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_listClientsCmd = &cobra.Command{
	Use:   "list-clients",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_listClientsCmd).Standalone()

	help_actionCmd.AddCommand(help_action_listClientsCmd)
}
