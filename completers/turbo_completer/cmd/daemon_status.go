package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var daemon_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Reports the status of the turbo daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_statusCmd).Standalone()

	daemonCmd.AddCommand(daemon_statusCmd)
}
