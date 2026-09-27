package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_testCmd).Standalone()

	plugin_testCmd.Flags().String("asdf-plugin-gitref", "", "The plugin Git ref to test")
	plugin_testCmd.Flags().String("asdf-tool-version", "", "The tool version to use during testing")
	pluginCmd.AddCommand(plugin_testCmd)
}
