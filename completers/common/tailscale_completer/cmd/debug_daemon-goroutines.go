package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_daemon_goroutinesCmd = &cobra.Command{
	Use:   "daemon-goroutines",
	Short: "Print tailscaled's goroutines",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_daemon_goroutinesCmd).Standalone()

	debugCmd.AddCommand(debug_daemon_goroutinesCmd)
}
