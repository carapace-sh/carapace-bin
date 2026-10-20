package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate a plugin or marketplace manifest, or the skills, agents, and commands in a directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_validateCmd).Standalone()

	plugin_validateCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_validateCmd.Flags().Bool("json", false, "Output the validation report as JSON (same exit codes)")
	plugin_validateCmd.Flags().Bool("strict", false, "Treat warnings as errors (exit 1)")
	pluginCmd.AddCommand(plugin_validateCmd)
}
