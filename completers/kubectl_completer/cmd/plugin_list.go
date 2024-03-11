package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all visible plugin executables on a user's PATH",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_listCmd).Standalone()

	plugin_listCmd.Flags().Bool("name-only", false, "If true, display only the binary name of each plugin, rather than its full path")
	pluginCmd.AddCommand(plugin_listCmd)
}
