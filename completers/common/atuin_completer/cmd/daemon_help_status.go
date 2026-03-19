package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_help_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the daemon's current status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_help_statusCmd).Standalone()

	daemon_helpCmd.AddCommand(daemon_help_statusCmd)
}
