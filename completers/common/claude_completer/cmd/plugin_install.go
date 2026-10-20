package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install a plugin from available marketplaces (use plugin@marketplace for specific marketplace)",
	Aliases: []string{"i"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_installCmd).Standalone()

	plugin_installCmd.Flags().String("accept-command", "", "Accept the marketplace-declared command…")
	plugin_installCmd.Flags().String("config", "", "Set a userConfig option declared in the plugin's manifest, or a bundled .mcpb server's own…")
	plugin_installCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_installCmd.Flags().Bool("json", false, "Print one machine-readable result line on stdout instead of the human message…")
	plugin_installCmd.Flags().String("marketplace", "", "Install <plugin> (its bare name) from the marketplace at this source: owner/repo, a git or https…")
	plugin_installCmd.Flags().String("registry", "", "For a <package>@npm install: resolve and download from this npm registry instead of the one your…")
	plugin_installCmd.Flags().StringP("scope", "s", "", "Installation scope: user, project, or local")
	plugin_installCmd.Flags().BoolP("yes", "y", false, "Accept the displayed marketplace-declared command without the confirmation prompt — a plugin…")
	pluginCmd.AddCommand(plugin_installCmd)

	carapace.Gen(plugin_installCmd).FlagCompletion(carapace.ActionMap{
		"scope": carapace.ActionValues("user", "project", "local"),
	})
}
