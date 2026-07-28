package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_daemonLogsCmd = &cobra.Command{
	Use:   "daemon-logs",
	Short: "Watch tailscaled's server logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_daemonLogsCmd).Standalone()

	debug_daemonLogsCmd.Flags().Bool("time", false, "include client time")
	debug_daemonLogsCmd.Flags().Int("verbose", 0, "verbosity level")
	debugCmd.AddCommand(debug_daemonLogsCmd)
}
