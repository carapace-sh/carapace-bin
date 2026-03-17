package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_daemon_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the daemon server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_daemon_startCmd).Standalone()

	help_daemonCmd.AddCommand(help_daemon_startCmd)
}
