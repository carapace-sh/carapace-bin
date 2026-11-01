package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverLogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Show detached server logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverLogsCmd).Standalone()
	serverLogsCmd.Flags().Bool("all", false, "show the entire available detached log")
	serverLogsCmd.Flags().Bool("follow", false, "follow appended log output until the server stops")
	serverLogsCmd.Flags().StringP("lines", "n", "", "show the newest N log lines")
	serverLogsCmd.Flags().String("state-dir", "", "read server logs below DIR")

	serverCmd.AddCommand(serverLogsCmd)

	carapace.Gen(serverLogsCmd).FlagCompletion(carapace.ActionMap{
		"state-dir": carapace.ActionDirectories(),
	})
}
