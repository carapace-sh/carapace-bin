package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_channelCmd = &cobra.Command{
	Use:   "channel",
	Short: "Commands to manage workspace channels",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_channelCmd).Standalone()

	workspaceCmd.AddCommand(workspace_channelCmd)
}
