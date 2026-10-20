package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a plugin to the latest version (restart required to apply)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_updateCmd).Standalone()

	plugin_updateCmd.Flags().String("accept-command", "", "Accept the marketplace-declared command…")
	plugin_updateCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_updateCmd.Flags().Bool("json", false, "Print one machine-readable result line on stdout instead of the human message…")
	plugin_updateCmd.Flags().StringP("scope", "s", "", "Installation scope: user, project, local, managed")
	plugin_updateCmd.Flags().BoolP("yes", "y", false, "Accept the displayed marketplace-declared command without the confirmation prompt — a changed…")
	pluginCmd.AddCommand(plugin_updateCmd)

	carapace.Gen(plugin_updateCmd).FlagCompletion(carapace.ActionMap{
		"scope": carapace.ActionValues("user", "project", "local", "managed"),
	})
}
