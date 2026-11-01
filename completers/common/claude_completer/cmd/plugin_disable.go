package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable an enabled plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_disableCmd).Standalone()

	plugin_disableCmd.Flags().BoolP("all", "a", false, "Disable all enabled plugins")
	plugin_disableCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_disableCmd.Flags().Bool("json", false, "Print one machine-readable result line on stdout instead of the human message (same exit codes)")
	plugin_disableCmd.Flags().StringP("scope", "s", "", "Installation scope: user, project, local")
	pluginCmd.AddCommand(plugin_disableCmd)

	carapace.Gen(plugin_disableCmd).FlagCompletion(carapace.ActionMap{
		"scope": carapace.ActionValues("local", "project", "user"),
	})
}
