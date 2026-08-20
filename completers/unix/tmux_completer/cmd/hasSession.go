package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tmux"
	"github.com/spf13/cobra"
)

var hasSessionCmd = &cobra.Command{
	Use:     "has-session",
	Aliases: []string{"has"},
	Short:   "check and report if a session exists on the server",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hasSessionCmd).Standalone()

	hasSessionCmd.Flags().StringS("t", "t", "", "specify target session")
	rootCmd.AddCommand(hasSessionCmd)

	carapace.Gen(hasSessionCmd).FlagCompletion(carapace.ActionMap{
		"t": tmux.ActionSessions(),
	})
}
