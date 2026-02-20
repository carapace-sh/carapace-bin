package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_channel_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the channels in the manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_channel_listCmd).Standalone()

	workspace_help_channelCmd.AddCommand(workspace_help_channel_listCmd)
}
