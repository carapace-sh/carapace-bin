package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_channel_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a channel to the manifest and updates the lockfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_channel_addCmd).Standalone()

	help_workspace_channelCmd.AddCommand(help_workspace_channel_addCmd)
}
