package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_channel_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List the channels in the manifest",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_channel_listCmd).Standalone()

	workspace_channel_listCmd.Flags().Bool("urls", false, "Whether to display the channel's names or urls")
	workspace_channelCmd.AddCommand(workspace_channel_listCmd)
}
