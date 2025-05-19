package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restarts the turbo daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_restartCmd).Standalone()

	daemonCmd.AddCommand(daemon_restartCmd)
}
