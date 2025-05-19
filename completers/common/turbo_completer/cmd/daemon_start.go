package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Ensures that the turbo daemon is running",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_startCmd).Standalone()

	daemonCmd.AddCommand(daemon_startCmd)
}
