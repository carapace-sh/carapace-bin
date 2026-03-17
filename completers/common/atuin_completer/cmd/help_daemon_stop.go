package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_daemon_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the daemon gracefully",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_daemon_stopCmd).Standalone()

	help_daemonCmd.AddCommand(help_daemon_stopCmd)
}
