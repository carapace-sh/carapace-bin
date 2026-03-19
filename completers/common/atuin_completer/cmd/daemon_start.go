package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the daemon server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_startCmd).Standalone()

	daemon_startCmd.Flags().Bool("force", false, "Force start: kill existing daemon process and reset the socket")
	daemon_startCmd.Flags().BoolP("help", "h", false, "Print help")
	daemon_startCmd.Flags().Bool("show-logs", false, "Also write daemon logs to the console (useful for debugging)")
	daemonCmd.AddCommand(daemon_startCmd)
}
