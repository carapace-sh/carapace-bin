package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_channel_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the channels in the manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_channel_help_listCmd).Standalone()

	workspace_channel_helpCmd.AddCommand(workspace_channel_help_listCmd)
}
