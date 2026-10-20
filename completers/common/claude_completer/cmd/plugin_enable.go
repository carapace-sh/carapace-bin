package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a disabled plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_enableCmd).Standalone()

	plugin_enableCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_enableCmd.Flags().Bool("json", false, "Print one machine-readable result line on stdout instead of the human message (same exit codes)")
	plugin_enableCmd.Flags().StringP("scope", "s", "", "Installation scope: user, project, local")
	pluginCmd.AddCommand(plugin_enableCmd)

	carapace.Gen(plugin_enableCmd).FlagCompletion(carapace.ActionMap{
		"scope": carapace.ActionValues("user", "project", "local"),
	})
}
