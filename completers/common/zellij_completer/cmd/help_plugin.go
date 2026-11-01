package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Load a plugin Returns: Created pane ID (format: plugin_<id>)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_pluginCmd).Standalone()

	helpCmd.AddCommand(help_pluginCmd)
}
