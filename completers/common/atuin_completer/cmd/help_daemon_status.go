package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_daemon_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the daemon's current status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_daemon_statusCmd).Standalone()

	help_daemonCmd.AddCommand(help_daemon_statusCmd)
}
