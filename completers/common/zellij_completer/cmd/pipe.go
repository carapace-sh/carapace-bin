package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipeCmd = &cobra.Command{
	Use:   "pipe",
	Short: "Send data to one or more plugins, launch them if they are not running",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipeCmd).Standalone()

	pipeCmd.Flags().StringP("args", "a", "", "The args of the pipe")
	pipeCmd.Flags().BoolP("help", "h", false, "Print help")
	pipeCmd.Flags().StringP("name", "n", "", "The name of the pipe")
	pipeCmd.Flags().StringP("plugin", "p", "", "The plugin url (eg. file:/tmp/my-plugin.wasm) to direct this pipe to, if not specified, will be sent to all plugins, if specified and is not running, the plugin will be launched")
	pipeCmd.Flags().StringP("plugin-configuration", "c", "", "The plugin configuration (note: the same plugin with different configuration is considered a different plugin for the purposes of determining the pipe destination)")
	rootCmd.AddCommand(pipeCmd)
}
