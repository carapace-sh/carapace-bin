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
	workspace_channelCmd.Flags().String("config-file", "", "Load configuration from this file instead of searching system and user-level paths. Project-local `<project>/.pixi/config.toml` is still merged on top")
	workspace_channelCmd.Flags().Bool("no-config", false, "Don't read system or user-level configuration files. Project-local `<project>/.pixi/config.toml` is still loaded")
	workspace_channelCmd.PersistentFlags().StringP("script", "s", "", "The path to a script with an embedded manifest: PEP 723 Python metadata or a conda-script block. Pixi run also accepts an HTTP(S) URL or '-' for stdin")

	workspaceCmd.AddCommand(workspace_channelCmd)

	carapace.Gen(workspace_channelCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
		"script":      carapace.ActionFiles(),
	})
}
