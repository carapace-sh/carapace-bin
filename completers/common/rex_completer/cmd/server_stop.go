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

	serverStopCmd.Flags().Bool("discard", false, "delete the saved sessions once the server has stopped")
	serverStopCmd.Flags().Bool("if-idle", false, "stop only if the server has no sessions and no other connected clients")
	serverStopCmd.Flags().String("state-dir", "", "find the saved sessions for --discard below DIR")

	serverCmd.AddCommand(serverStopCmd)

	carapace.Gen(serverStopCmd).FlagCompletion(carapace.ActionMap{
		"state-dir": carapace.ActionDirectories(),
	})
}
