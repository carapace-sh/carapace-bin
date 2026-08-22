package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_pipeCmd = &cobra.Command{
	Use:   "pipe",
	Short: "Send data to one or more plugins, launch them if they are not running",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_pipeCmd).Standalone()

	action_pipeCmd.Flags().StringP("args", "a", "", "The args of the pipe")
	action_pipeCmd.Flags().StringP("floating-plugin", "f", "", "If launching a plugin, should it be floating or not, defaults to floating")
	action_pipeCmd.Flags().BoolP("force-launch-plugin", "l", false, "Launch a new plugin even if one is already running")
	action_pipeCmd.Flags().BoolP("help", "h", false, "Print help")
	action_pipeCmd.Flags().StringP("in-place-plugin", "i", "", "If launching a plugin, launch it in-place (on top of the current pane)")
	action_pipeCmd.Flags().StringP("name", "n", "", "The name of the pipe")
	action_pipeCmd.Flags().StringP("plugin", "p", "", "The plugin url (eg. file:/tmp/my-plugin.wasm) to direct this pipe to, if not specified, will be sent to all plugins, if specified and is not running, the plugin will be launched")
	action_pipeCmd.Flags().StringP("plugin-configuration", "c", "", "The plugin configuration (note: the same plugin with different configuration is considered a different plugin for the purposes of determining the pipe destination)")
	action_pipeCmd.Flags().StringP("plugin-cwd", "w", "", "If launching a plugin, specify its working directory")
	action_pipeCmd.Flags().StringP("plugin-title", "t", "", "If launching a plugin, specify its pane title")
	action_pipeCmd.Flags().BoolP("skip-plugin-cache", "s", false, "If launching a new plugin, skip cache and force-compile the plugin")
	actionCmd.AddCommand(action_pipeCmd)

	carapace.Gen(action_pipeCmd).FlagCompletion(carapace.ActionMap{
		"floating-plugin": carapace.ActionValues("true", "false"),
		"in-place-plugin": carapace.ActionValues("true", "false"),
		"plugin-cwd":      carapace.ActionDirectories(),
	})
}
