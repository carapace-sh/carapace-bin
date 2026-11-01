package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Remove the background service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_uninstallCmd).Standalone()

	daemonCmd.AddCommand(daemon_uninstallCmd)
}
