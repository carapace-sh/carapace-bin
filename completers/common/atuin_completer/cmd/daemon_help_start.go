package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_help_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the daemon server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_help_startCmd).Standalone()

	daemon_helpCmd.AddCommand(daemon_help_startCmd)
}
