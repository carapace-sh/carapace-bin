package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var suspendClientCmd = &cobra.Command{
	Use:     "suspend-client",
	Aliases: []string{"suspendc"},
	Short:   "suspend a client",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspendClientCmd).Standalone()

	suspendClientCmd.Flags().StringS("t", "t", "", "specify target client")
	rootCmd.AddCommand(suspendClientCmd)

	carapace.Gen(suspendClientCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionClients(),
	})
}
