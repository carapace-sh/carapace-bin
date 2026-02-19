package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_channel_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the channels in the manifest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_channel_listCmd).Standalone()

	workspace_channel_listCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspace_channel_listCmd.Flags().Bool("urls", false, "Show the URLs of the channels")
	workspace_channelCmd.AddCommand(workspace_channel_listCmd)

	carapace.Gen(workspace_channel_listCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
