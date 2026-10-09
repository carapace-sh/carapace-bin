package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var lockClientCmd = &cobra.Command{
	Use:     "lock-client",
	Aliases: []string{"lockc"},
	Short:   "lock a client",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lockClientCmd).Standalone()

	lockClientCmd.Flags().StringS("t", "t", "", "specify target client")
	rootCmd.AddCommand(lockClientCmd)

	carapace.Gen(lockClientCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionClients(),
	})
}
