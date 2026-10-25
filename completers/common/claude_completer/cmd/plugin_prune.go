package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_pruneCmd = &cobra.Command{
	Use:     "prune",
	Short:   "Remove auto-installed dependencies that are no longer needed",
	Aliases: []string{"autoremove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_pruneCmd).Standalone()

	plugin_pruneCmd.Flags().Bool("dry-run", false, "List what would be removed without removing")
	plugin_pruneCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_pruneCmd.Flags().StringP("scope", "s", "", "Prune at scope: user, project, or local")
	plugin_pruneCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt (required when stdin or stdout is not a TTY)")
	pluginCmd.AddCommand(plugin_pruneCmd)

	carapace.Gen(plugin_pruneCmd).FlagCompletion(carapace.ActionMap{
		"scope": carapace.ActionValues("local", "project", "user"),
	})
}
