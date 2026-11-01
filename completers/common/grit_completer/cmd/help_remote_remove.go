package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_remote_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a remote and its remote-tracking refs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_remote_removeCmd).Standalone()

	help_remoteCmd.AddCommand(help_remote_removeCmd)
}
