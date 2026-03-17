package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_daemon_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart the daemon (stop, then start in background)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_daemon_restartCmd).Standalone()

	help_daemonCmd.AddCommand(help_daemon_restartCmd)
}
