package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var showMessagesCmd = &cobra.Command{
	Use:     "show-messages",
	Aliases: []string{"showmsgs"},
	Short:   "show client's message log",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showMessagesCmd).Standalone()

	showMessagesCmd.Flags().BoolS("J", "J", false, "show debugging information about running jobs")
	showMessagesCmd.Flags().BoolS("T", "T", false, "show debugging information about involved terminals")
	showMessagesCmd.Flags().StringS("t", "t", "", "specify target client")
	rootCmd.AddCommand(showMessagesCmd)

	carapace.Gen(showMessagesCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionClients(),
	})
}
