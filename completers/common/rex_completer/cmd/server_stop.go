package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverStopCmd).Standalone()

	serverCmd.AddCommand(serverStopCmd)
}
