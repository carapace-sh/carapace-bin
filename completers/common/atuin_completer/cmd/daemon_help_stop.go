package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_help_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the daemon gracefully",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_help_stopCmd).Standalone()

	daemon_helpCmd.AddCommand(daemon_help_stopCmd)
}
