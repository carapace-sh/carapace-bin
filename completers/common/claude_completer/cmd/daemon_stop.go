package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the supervisor and the sessions it hosts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_stopCmd).Standalone()

	daemon_stopCmd.Flags().Bool("any", false, "Also stop an on-demand (non-service) supervisor")
	daemon_stopCmd.Flags().Bool("keep-workers", false, "Leave background sessions running")
	daemonCmd.AddCommand(daemon_stopCmd)
}
