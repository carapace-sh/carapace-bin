package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_listClientsCmd = &cobra.Command{
	Use:   "list-clients",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_listClientsCmd).Standalone()

	action_listClientsCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_listClientsCmd)
}
