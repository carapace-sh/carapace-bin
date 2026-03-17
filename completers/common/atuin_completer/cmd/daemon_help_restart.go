package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_help_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the daemon (stop, then start in background)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_help_restartCmd).Standalone()

	daemon_helpCmd.AddCommand(daemon_help_restartCmd)
}
