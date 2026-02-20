package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_channel_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the channels in the manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_channel_listCmd).Standalone()

	help_workspace_channelCmd.AddCommand(help_workspace_channel_listCmd)
}
