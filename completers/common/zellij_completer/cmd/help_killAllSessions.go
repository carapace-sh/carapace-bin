package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_killAllSessionsCmd = &cobra.Command{
	Use:   "kill-all-sessions",
	Short: "Kill all sessions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_killAllSessionsCmd).Standalone()

	helpCmd.AddCommand(help_killAllSessionsCmd)
}
