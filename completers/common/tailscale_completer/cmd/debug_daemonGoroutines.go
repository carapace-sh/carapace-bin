package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_daemonGoroutinesCmd = &cobra.Command{
	Use:   "daemon-goroutines",
	Short: "Print tailscaled's goroutines",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_daemonGoroutinesCmd).Standalone()

	debugCmd.AddCommand(debug_daemonGoroutinesCmd)
}
