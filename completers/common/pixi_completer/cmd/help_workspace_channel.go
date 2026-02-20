package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_channelCmd = &cobra.Command{
	Use:   "channel",
	Short: "Commands to manage workspace channels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_channelCmd).Standalone()

	help_workspaceCmd.AddCommand(help_workspace_channelCmd)
}
