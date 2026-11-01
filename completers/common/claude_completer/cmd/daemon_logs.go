package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Follow the supervisor log",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_logsCmd).Standalone()

	daemonCmd.AddCommand(daemon_logsCmd)
}
