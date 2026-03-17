package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the daemon gracefully",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_stopCmd).Standalone()

	daemon_stopCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonCmd.AddCommand(daemon_stopCmd)
}
