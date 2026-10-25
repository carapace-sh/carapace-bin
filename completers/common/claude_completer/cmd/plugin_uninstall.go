package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_uninstallCmd = &cobra.Command{
	Use:     "uninstall",
	Short:   "Uninstall an installed plugin",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_uninstallCmd).Standalone()

	plugin_uninstallCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_uninstallCmd.Flags().Bool("json", false, "Print one machine-readable result line on stdout instead of the human message…")
	plugin_uninstallCmd.Flags().Bool("keep-data", false, "Preserve the plugin's persistent data directory (~/.claude/plugins/data/{id}/)")
	plugin_uninstallCmd.Flags().Bool("prune", false, "Also remove auto-installed dependencies that are no longer needed…")
	plugin_uninstallCmd.Flags().StringP("scope", "s", "", "Uninstall from scope: user, project, or local")
	plugin_uninstallCmd.Flags().BoolP("yes", "y", false, "Skip the --prune confirmation prompt (required when stdin or stdout is not a TTY)")
	pluginCmd.AddCommand(plugin_uninstallCmd)

	carapace.Gen(plugin_uninstallCmd).FlagCompletion(carapace.ActionMap{
		"scope": carapace.ActionValues("local", "project", "user"),
	})
}
