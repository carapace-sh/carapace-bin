package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the turbo daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_stopCmd).Standalone()

	daemonCmd.AddCommand(daemon_stopCmd)
}
