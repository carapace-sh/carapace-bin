package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var forwardCloseCmd = &cobra.Command{
	Use:   "close [FORWARD...]",
	Short: "Close forwards",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forwardCloseCmd).Standalone()

	forwardCloseCmd.Flags().Bool("all", false, "close every forward in the session")
	forwardCloseCmd.Flags().Bool("orphaned", false, "close every orphaned forward in the session")
	forwardCloseCmd.Flags().StringP("session", "s", "", "scope to session SESSION")

	forwardCmd.AddCommand(forwardCloseCmd)

	carapace.Gen(forwardCloseCmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(forwardCloseCmd),
	})

	carapace.Gen(forwardCloseCmd).PositionalAnyCompletion(
		action.ActionBlocks(forwardCloseCmd),
	)
}
