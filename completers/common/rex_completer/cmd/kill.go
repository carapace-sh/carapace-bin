package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill SESSION...",
	Short: "Kill sessions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killCmd).Standalone()

	rootCmd.AddCommand(killCmd)

	carapace.Gen(killCmd).PositionalAnyCompletion(
		action.ActionSessions(killCmd),
	)
}
