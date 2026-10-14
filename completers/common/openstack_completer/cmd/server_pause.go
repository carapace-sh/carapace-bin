package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var server_pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause server(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(server_pauseCmd).Standalone()

	serverCmd.AddCommand(server_pauseCmd)
}
