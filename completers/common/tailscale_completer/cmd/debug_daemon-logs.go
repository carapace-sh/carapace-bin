package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_daemon_logsCmd = &cobra.Command{
	Use:   "daemon-logs",
	Short: "Watch tailscaled's server logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_daemon_logsCmd).Standalone()

	debugCmd.AddCommand(debug_daemon_logsCmd)
}
