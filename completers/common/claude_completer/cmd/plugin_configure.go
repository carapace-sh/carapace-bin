package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Show a plugin's options and which are unset, or save values from stdin with --values-stdin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_configureCmd).Standalone()

	plugin_configureCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_configureCmd.Flags().Bool("json", false, "Output as JSON")
	plugin_configureCmd.Flags().Bool("values-stdin", false, "Read option values from stdin as a JSON object of single-line strings; options left out keep their…")
	pluginCmd.AddCommand(plugin_configureCmd)
}
