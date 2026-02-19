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

	workspace_channelCmd.Flags().StringP("manifest-path", "m", "", "The path to pixi.toml, pyproject.toml, or the workspace directory")
	workspaceCmd.AddCommand(workspace_channelCmd)

	carapace.Gen(workspace_channelCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
	})
}
