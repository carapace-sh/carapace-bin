package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var pluginCmd = &cobra.Command{
	Use:     "plugin",
	Short:   "Load a plugin Returns: Created pane ID (format: plugin_<id>)",
	Aliases: []string{"p"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginCmd).Standalone()

	pluginCmd.Flags().StringP("borderless", "b", "", "start this pane without a border (warning: will make it impossible to move with the mouse)")
	pluginCmd.Flags().Bool("close-replaced-pane", false, "Close the replaced pane instead of suspending it (only effective with --in-place)")
	pluginCmd.Flags().StringP("configuration", "c", "", "Plugin configuration")
	pluginCmd.Flags().BoolP("floating", "f", false, "Open the new pane in floating mode")
	pluginCmd.Flags().String("height", "", "The height if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	pluginCmd.Flags().BoolP("help", "h", false, "Print help")
	pluginCmd.Flags().BoolP("in-place", "i", false, "Open the new pane in place of the current pane, temporarily suspending it")
	pluginCmd.Flags().Bool("no-focus", false, "if set, will open the plugin pane without changing the focus of any client, placing it relative to the pane the command was issued from")
	pluginCmd.Flags().String("pinned", "", "Whether to pin a floating pane so that it is always on top")
	pluginCmd.Flags().BoolP("skip-plugin-cache", "s", false, "Skip the memory and HD cache and force recompile of the plugin (good for development)")
	pluginCmd.Flags().String("tab-id", "", "Target a specific tab by ID")
	pluginCmd.Flags().String("width", "", "The width if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	pluginCmd.Flags().StringP("x", "x", "", "The x coordinates if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	pluginCmd.Flags().StringP("y", "y", "", "The y coordinates if the pane is floating as a bare integer (eg. 1) or percent (eg. 10%)")
	rootCmd.AddCommand(pluginCmd)

	carapace.Gen(pluginCmd).FlagCompletion(carapace.ActionMap{
		"borderless": carapace.ActionValues("true", "false").StyleF(style.ForKeyword),
		"pinned":     carapace.ActionValues("true", "false").StyleF(style.ForKeyword),
		"tab-id":     zellij.ActionTabs(),
	})
}
