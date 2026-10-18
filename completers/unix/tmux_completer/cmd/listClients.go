package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var listClientsCmd = &cobra.Command{
	Use:     "list-clients",
	Aliases: []string{"lsc"},
	Short:   "list clients attached to server",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listClientsCmd).Standalone()

	listClientsCmd.Flags().StringS("F", "F", "", "specify output format")
	listClientsCmd.Flags().StringS("O", "O", "", "initial sort order")
	listClientsCmd.Flags().StringS("f", "f", "", "filter items")
	listClientsCmd.Flags().BoolS("r", "r", false, "reverse sort order")
	listClientsCmd.Flags().StringS("t", "t", "", "specify target session")
	rootCmd.AddCommand(listClientsCmd)

	carapace.Gen(listClientsCmd).FlagCompletion(carapace.ActionMap{
		"O": carapace.ActionValues("name", "size", "creation", "activity"),
		"t": tmux.ActionSessions(),
	})
}
